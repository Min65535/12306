package web

import (
	"log"
	"strings"
	"time"
	"flag"
	"me.org/http"
	"net/url"
	"errors"
	"strconv"
	"12306/captcha"
)

/*
 * SingerOrder
 *
 * 提交单程订单
 *
 * @param Username 				string			用户名
 * @param Password 				string			密码
 * @param fromStation 			[]string		出发站信息,例:["深圳","SZQ"]
 * @param toStation 			[]string		目的站信息,例:["武汉","WHN"]
 * @param fromDate 				string			购票车次日期,例:"2017-12-10"
 * @param toDate 				string			回程日期,例:"2017-12-10";另需注意:单程此时应该传递空字段:""
 * @param tourFlag 				string			行程类型,例:dc,表示单程;wc,表示往返程
 * @param personType 			string			乘客类型,例:"ADULT"
 * @param stationTrainCode 		string			列车车次.例:"K238"
 * @param seatType 				string			座位类型.例:3,硬卧;4,软卧;0,二等座;M,一等座;1,硬座;
 *
 * @return bool		成功标识
 * @return error
 *
 */
func SingerOrder(Username string, Password string, fromStation []string, toStation []string, fromDate string, toDate string, tourFlag string, personType string, stationTrainCode string, passengers []string) {
	for _, arg := range flag.Args() {
		if strings.ContainsAny(arg, "=") {
			_arg := strings.Split(arg, "=")
			if _arg[0] == "username" {
				Username = _arg[1]
			} else if _arg[0] == "password" {
				Password = _arg[1]
			}
		}
	}
	if Username == "" || Password == "" {
		log.Fatalln("please input 12306 username and password in Run -> Edit Configurations -> [Go Test] -> TestLogin in login_test.go -> Program arguments set like username=用户名 password=用户密码")
	}

	log.Println("12306 username:", Username, "Password:", Password)

	client := http.NewClient(nil, "", "", 0)
	client.SetInsecureSkipVerify(true)
	var sign = false               //错误标识
	var errorName string           //错误名字
	var totalError error           //错误原因
	var secretStr string           //从LeftTicketQuery取出的行程密钥
	var TicketResult []string      //从LeftTicketQuery取出的行程数组
	var trainDate string           //乘车日期//Sun+Dec+10+2017+00:00:00+GMT+0800
	var trainNo string             //列车编号//690000K2380A
	var fromStationTelecode string //始发站的三字码
	var toStationTelecode string   //终点站的三字码
	var leftTicket string          //LeftTicketQuery里的余票短密钥,例:Wey1KErtCNm8%2B1Bt55Vy9F9hqwuUR0KY%2F2P5DarryhPxyPnkDbWGDEP775g%3D
	var trainLocation string       //LeftTicketQuery里的短密钥后面第三段里的Q7        //|20171210|3|Q7|01|11|0|0||||10|||有||有|有|||||10401030|1413
	var orderId string             //列车单号
	var requestId int              //列车购票请求ID
	for i := 0; i < JudgeNum; i++ {
		for i := 0; i < JudgeNum; i++ {
			cookies, err := Init(client, InitUrl)
			if err != nil {
				errorName = "Init"
				totalError = err
			} else {
				log.Println("Init new cookies:", cookies)
				result, cookies, err := Uamtk(client, UamtkUrl, nil)
				if err != nil {
					errorName = "Uamtk"
					totalError = err
				} else {
					log.Println("Uamtk result:", result, "new cookies:", cookies)
					log.Println("client cookie:", client.Cookie())

					myCaptcha := captcha.GetCaptcha()

					for i := 0; i < JudgeNum; i++ {
						// 获取验证码
						image, cookies, err := GetCaptcha(client, CaptchaUrl, nil)
						if err != nil {
							errorName = "GetCaptcha"
							totalError = err
							continue
						}

						//log.Println("GetCaptcha image length:", len(image), "new cookies:", cookies)

						log.Println("client cookie:", client.Cookie())

						// 远程打码
						id, value, err := myCaptcha.Upload("Jsdama-12306", image)
						if err != nil {
							errorName = "Upload"
							totalError = err
							continue
						}

						//log.Println("Upload CaptchaId:", id, "Captcha:", value)

						// 检查验证码
						result, cookies, succeed, err := CheckCaptcha(client, CaptchaCheckUrl, nil, value)
						if err != nil {
							errorName = "CheckCaptcha"
							totalError = err
							continue
						}

						log.Println("CheckCaptcha result:", result, "succeed:", succeed, "new cookie:", cookies)

						if !succeed {
							myCaptcha.Report("Jsdama-12306", id)
							log.Println("sleep 2s")
							time.Sleep(2 * time.Second)
							continue
						}

						result, cookies, succeed, err = Login(client, LoginUrl, nil, Username, Password)
						if err != nil {
							errorName = "Login"
							totalError = err
							continue
						}

						log.Println("Login result:", result, "succeed:", succeed, "new cookies:", cookies)

						log.Println("client cookie:", client.Cookie())

						if succeed {
							cookies, err = UserLogin(client, UserLoginUrl, nil)
							if err != nil {
								errorName = "UserLogin"
								totalError = err
							} else {
								log.Println("UserLogin new cookies:", cookies)
								log.Println("client cookie:", client.Cookie())
								result, cookies, err = UamtkLogin(client, UamtkUrl, nil)
								if err != nil {
									errorName = "UamtkLogin"
									totalError = err
								} else {

									log.Println("UamtkLogin result:", result, "new cookies:", cookies)

									log.Println("client cookie:", client.Cookie())

									result, cookies, err = UamauthClient(client, UamauthClientUrl, nil, result.GetTK())
									if err != nil {
										errorName = "UamauthClient"
										totalError = err
									} else {
										if result.Username == "" {
											errorName = "user's name obtain"
											totalError = errors.New("fail to obtain the user's name")
										} else {
											log.Println("UamauthClient result:", result, "new cookies:", cookies)

											log.Println("client cookie:", client.Cookie())

											log.Println("welcome to", result.Username)

											_, err := UserLoginSec(client, UserLoginSecUrl, nil)
											if err != nil {
												errorName = "UserLoginSec"
												totalError = err
											} else {
												_, err := InitMy12306(client, InitMy12306Url, nil)
												if err != nil {
													errorName = "InitMy12306"
													totalError = err
												} else {
													sign = true
												}
											}

										}
									}
								}
								break
							}
						}
					}
					if sign == true {
						sign = false
						_, err := LeftTicketInit(client, LeftTicketInitUrl, nil)
						if err != nil {
							errorName = "LeftTicketInit"
							totalError = err
						} else {
							_, err := GetPassCodeNew(client, GetPassCodeNewUrl, nil)
							if err != nil {
								errorName = "GetPassCodeNew"
								totalError = err
							} else {
								for i := 0; i < JudgeNum; i++ {
									logSign, err := LeftTicketLog(client, LeftTicketLogUrl, nil, fromStation, toStation, fromDate, toDate, tourFlag, personType)
									if err != nil {
										errorName = "LeftTicketLog"
										totalError = err
									} else {
										if logSign != true {
											errorName = "LeftTicketLog"
											totalError = errors.New("LeftTicketLog get an incorrect response")
										} else {
											msgJson, err := LeftTicketQuery(client, LeftTicketQueryUrl, nil, fromStation, toStation, fromDate, personType)
											if err != nil {
												errorName = "LeftTicketQuery"
												totalError = err
											} else {
												log.Println("LeftTicketQuery result:", msgJson)
												log.Println("client cookie:", client.Cookie())

												TicketResult = msgJson.Data.Result

												log.Println("TicketResult value:", TicketResult)

												if len(TicketResult) == 0 {
													errorName = "LeftTicketQuery"
													totalError = errors.New("the data result of LeftTicketQuery is empty")
												}
												if len(TicketResult) != 0 {
													sign = true
												}
											}

										}
									}
									if sign == true {
										break
									}
									log.Println("LeftTicket Search occur error,sleep 3s")
									time.Sleep(3 * time.Second)
								}
								if sign == true {
									sign = false
									CheckSign, err := CheckUser(client, CheckUserUrl, nil)
									log.Println("CheckUser sign:", CheckSign)
									log.Println("client cookie:", client.Cookie())
									if err != nil {
										errorName = "CheckUser"
										totalError = err
									} else {
										if CheckSign == false {
											errorName = "CheckUser response"
											totalError = errors.New("CheckUser get an incorrect response")
										} else {
											sign = CheckSign
											break
										}
										log.Println("CheckUser occur error,sleep 2s")
										time.Sleep(2 * time.Second)
									}
								}
							}
						}
					}
				}
			}
			if sign == true {
				break
			}
		}
		if sign == true {
			sign = false
			for _, b := range TicketResult {
				if strings.Contains(b, stationTrainCode) {
					arr := strings.Split(b, "|")
					log.Println("True TicketResult array:", arr)
					str := arr[0]
					temptStr := strings.Replace(str, `%0A`, ``, -1)
					ww, _ := url.QueryUnescape(temptStr)
					secretStr = ww

					fromDateTemp := fromDate + ` 00:00:00`
					date, _ := time.Parse("2006-01-02 15:04:05", fromDateTemp)
					weekStr := date.Weekday().String()
					monthStr := date.Month().String()
					dayStr := strconv.Itoa(date.Day())
					yearStr := strconv.Itoa(date.Year())
					trainDate = weekStr[:3] + ` ` + monthStr[:3] + ` ` + dayStr + ` ` + yearStr + ` 00:00:00 GMT+0800 (中国标准时间)`

					trainNo = arr[2]
					fromStationTelecode = arr[6]
					toStationTelecode = arr[7]

					leftTicket = arr[12]

					trainLocation = arr[15]
					break
				}
			}
			if secretStr == "" {
				errorName = "get secretStr"
				totalError = errors.New("fail to get secretStr")
			} else {
				log.Println("secretStr value:", secretStr)
				for i := 0; i < JudgeNum; i++ {
					submitSign, err := SubmitOrderRequest(client, SubmitOrderRequestUrl, nil, secretStr, fromDate, toDate, tourFlag, personType, fromStation[0], toStation[0])
					if err != nil {
						errorName = "SubmitOrderRequest"
						totalError = err
					} else {
						log.Println("SubmitOrderRequest sign:", submitSign)
						if submitSign == false {
							errorName = "SubmitOrderRequest"
							totalError = errors.New("SubmitOrderRequest get an incorrect response")
						}
						if submitSign == true {
							sign = submitSign
							break
						}
					}
				}
			}
		}
		if sign == true {
			sign = false
			var repeatSubmitToken string
			var keyCheckIsChange string
			repeatSubmitToken, keyCheckIsChange, err := InitDc(client, InitDcUrl, nil)
			if err != nil {
				errorName = "InitDc"
				totalError = err
			} else {
				log.Println("repeatSubmitToken value:", repeatSubmitToken)
				log.Println("keyCheckIsChange value:", keyCheckIsChange)
				if repeatSubmitToken == "" || keyCheckIsChange == "" {
					errorName = "InitDc"
					totalError = errors.New("fail to get repeatSubmitToken or keyCheckIsChange")
				} else {
					passengerJson, err := GetPassengerDTOs(client, GetPassengerDTOsUrl, nil, repeatSubmitToken)
					if err != nil {
						errorName = "GetPassengerDTOs"
						totalError = err
					} else {
						log.Println("GetPassengerDTOs value:", passengerJson)
						log.Println("NormalPassengers value:", passengerJson.Data.NormalPassengers)
						if passengers == nil {
							errorName = "passengers"
							totalError = errors.New("the value of passengers is empty")
						} else {
							normalPassengers := passengerJson.Data.NormalPassengers
							var passengerTicketStr string
							var passengerTicketStrArr []string
							var passengerTicketStrTemp string
							var oldPassengerStr string
							var oldPassengerStrArr []string
							var oldPassengerStrTemp string
							var seatType string
							seatType = (strings.Split(passengers[0], `,`))[0]
							log.Println("seatType value:", seatType)
							for i := 0; i < len(passengers); i++ {
								msgArr := strings.Split(passengers[i], `,`)
								trainType := msgArr[0]
								ticketType := msgArr[1]
								passengerName := msgArr[2]
								for j := 0; j < len(normalPassengers); j++ {
									if normalPassengers[j].PassengerName == passengerName {
										passengerTicketStrArr = append(passengerTicketStrArr, trainType+`,0,`+ticketType+`,`+passengerName+`,`+normalPassengers[j].PassengerIdTypeCode+`,`+normalPassengers[j].PassengerIdNo+`,`+normalPassengers[j].MobileNo+`,N_`)
										if ticketType == "2" {
											oldPassengerStrArr = append(oldPassengerStrArr, "_+")
										} else if ticketType == "1" {
											oldPassengerStrArr = append(oldPassengerStrArr, passengerName+`,`+normalPassengers[j].PassengerIdTypeCode+`,`+normalPassengers[j].PassengerIdNo+`,`+normalPassengers[j].PassengerType+`_`)
										}
									}
								}
							}

							log.Println("passengerTicketStrArr value:", passengerTicketStrArr)
							log.Println("oldPassengerStrArr value:", oldPassengerStrArr)
							for x := 0; x < len(passengerTicketStrArr); x++ {
								passengerTicketStrTemp += passengerTicketStrArr[x]
								oldPassengerStrTemp += oldPassengerStrArr[x]
							}
							passengerTicketStr = passengerTicketStrTemp[0:len(passengerTicketStrTemp)-1]
							oldPassengerStr = oldPassengerStrTemp
							log.Println("passengerTicketStr value:", passengerTicketStr)
							log.Println("oldPassengerStr value:", oldPassengerStr)
							_, err := GetPassCodeNewSubmit(client, GetPassCodeNewUrl, nil)
							if err != nil {
								errorName = "GetPassCodeNewSubmit"
								totalError = err
							} else {
								checkOrderJson, err := CheckOrderInfo(client, CheckOrderInfoUrl, nil, passengerTicketStr, oldPassengerStr, tourFlag, repeatSubmitToken)
								if err != nil {
									errorName = "CheckOrderInfo"
									totalError = err
								} else {
									log.Println("CheckOrderInfo value:", checkOrderJson)
									log.Println("CheckOrderJson.Status value:", checkOrderJson.Status)
									log.Println("CheckOrderJson.Data.IfShowPassCode value:", checkOrderJson.Data.IfShowPassCode)
									if checkOrderJson.Status != true {
										errorName = "CheckOrderInfo"
										totalError = errors.New("the Status of CheckOrderInfo is incorrect")
									} else {
										if checkOrderJson.Data.IfShowPassCode != "N" {
											errorName = "CheckOrderInfo"
											totalError = errors.New("the next step of CheckOrderInfo need captcha")
										} else {
											log.Println("trainDate value:", trainDate)
											log.Println("trainNo value:", trainNo)
											log.Println("stationTrainCode value:", stationTrainCode)
											log.Println("seatType value:", seatType)
											log.Println("fromStationTelecode value:", fromStationTelecode)
											log.Println("toStationTelecode value:", toStationTelecode)
											log.Println("leftTicket value:", leftTicket)
											log.Println("trainLocation value:", trainLocation)
											log.Println("repeatSubmitToken value:", repeatSubmitToken)

											getQueueCountJson, err := GetQueueCount(client, GetQueueCountUrl, nil, trainDate, trainNo, stationTrainCode, seatType, fromStationTelecode, toStationTelecode, leftTicket, trainLocation, repeatSubmitToken)
											if err != nil {
												errorName = "GetQueueCount"
												totalError = err
											} else {
												log.Println("GetQueueCount value:", getQueueCountJson)
												log.Println("getQueueCountJson.Data.Ticket value:", getQueueCountJson.Data.Ticket)

												ticketNumStr := getQueueCountJson.Data.Ticket
												ticketNum, err := strconv.ParseInt(ticketNumStr, 10, 0)
												if ticketNumStr == "" {
													errorName = "GetQueueCount"
													totalError = errors.New("ticketNumStr is empty")
												} else {
													if err != nil {
														errorName = "GetQueueCount"
														totalError = errors.New("ticketNumStr ParseInt occur error")
													} else {
														if ticketNum < 1 {
															errorName = "GetQueueCount"
															totalError = errors.New("the num of ticket is not enough")
														} else {
															confirmJson, err := ConfirmSingleForQueue(client, ConfirmSingleForQueueUrl, nil, passengerTicketStr, oldPassengerStr, keyCheckIsChange, leftTicket, trainLocation, repeatSubmitToken)
															if err != nil {
																errorName = "ConfirmSingleForQueue"
																totalError = err
															} else {
																log.Println("ConfirmSingleForQueue value:", confirmJson)
																if confirmJson.Status != true {
																	errorName = "ConfirmSingleForQueue"
																	totalError = errors.New("ConfirmSingleForQueue get an incorrect response")
																} else {
																	for i := 0; i < JudgeNum; i++ {
																		QueryOrderJson, err := QueryOrderWaitTime(client, QueryOrderWaitTimeUrl, nil, tourFlag, repeatSubmitToken)
																		if err != nil {
																			errorName = "ConfirmSingleForQueue"
																			totalError = err
																		} else {
																			log.Println("QueryOrderWaitTime value:", QueryOrderJson)

																			if QueryOrderJson.Data.OrderId == nil {
																				errorName = "ConfirmSingleForQueue"
																				totalError = errors.New("QueryOrderJson.Data.orderId is null")
																			} else {
																				orderId = QueryOrderJson.Data.OrderId.(string)
																				requestId = QueryOrderJson.Data.RequestId
																				sign = true
																			}
																		}
																		if sign == true {
																			break
																		}
																		log.Println("QueryOrderWaitTime go into the loop,sleep 4s")
																		time.Sleep(4 * time.Second)
																	}
																}
															}
														}
													}
												}
											}

										}
									}

								}
							}
						}
					}
				}
			}
		}

		if sign == true {
			break
		}

	}

	if sign == false {
		log.Println(errorName, " occur error:", totalError)
	} else {
		log.Println("Congratulations! you are lucky to get the ticket,the request ID is:", requestId, "and your ticket order ID is:", orderId)
	}

}
