//package poller
//
//import (
//	"github.com/chremoas/auth-srv/proto"
//	"github.com/chremoas/auth-srv/proto/mocks"
//	authsrv_matchers "github.com/chremoas/auth-srv/proto/mocks/matchers"
//	"github.com/chremoas/esi-srv/proto"
//	"github.com/chremoas/esi-srv/proto/mocks"
//	esi_matchers "github.com/chremoas/esi-srv/proto/mocks/matchers"
//	. "github.com/smartystreets/goconvey/convey"
//	"testing"
//	. "github.com/petergtz/pegomock"
//	"strings"
//	"regexp"
//	"runtime/debug"
//	"errors"
//)
//
//func PollerTestSetup_Pegomock(t *testing.T) (
//	*authsrv_mocks.MockEntityQueryClient,
//	*authsrv_mocks.MockEntityAdminClient,
//	*esi_mocks.MockAllianceServiceClient,
//	*esi_mocks.MockCorporationServiceClient,
//	*esi_mocks.MockCharacterServiceClient,
//) {
//	RegisterMockTestingT(t)
//	//RegisterMockFailHandler(BuildTestingTGomegaFailHandler(t))
//	RegisterMockFailHandler(NewFailHandler(t))
//
//	mockEntityQueryClient := authsrv_mocks.NewMockEntityQueryClient()
//	mockEntityAdminClient := authsrv_mocks.NewMockEntityAdminClient()
//	mockAllianceClient := esi_mocks.NewMockAllianceServiceClient()
//	mockCorporationClient := esi_mocks.NewMockCorporationServiceClient()
//	mockCharacterClient := esi_mocks.NewMockCharacterServiceClient()
//
//	return mockEntityQueryClient, mockEntityAdminClient, mockAllianceClient, mockCorporationClient, mockCharacterClient
//}
//
//func TestCompareFunctions(t *testing.T) {
//	Convey("Should be able to compare 2 alliances", t, func() {
//		authAlliance := &abaeve_auth.Alliance{
//			Id: 1,
//			Name: "Alliance Name 1",
//			Ticker: "A T 1",
//		}
//
//		Convey("Alliances differ by name", func() {
//			esiAlliance := &chremoas_esi.Alliance{
//				Name: "Alliance Name 2",
//				Ticker: "A T 1",
//			}
//
//			answer := allianceDiffers(authAlliance, esiAlliance)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Alliances differ by ticker", func() {
//			esiAlliance := &chremoas_esi.Alliance{
//				Name: "Alliance Name 1",
//				Ticker: "A T 2",
//			}
//
//			answer := allianceDiffers(authAlliance, esiAlliance)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Alliances are the same", func() {
//			esiAlliance := &chremoas_esi.Alliance{
//				Name: "Alliance Name 1",
//				Ticker: "A T 1",
//			}
//
//			answer := allianceDiffers(authAlliance, esiAlliance)
//
//			So(answer, ShouldEqual, false)
//		})
//	})
//
//	Convey("Should be able to compare 2 corporations", t, func() {
//		authCorporation := &abaeve_auth.Corporation{
//			Id: 1,
//			Name: "Corporation Name 1",
//			Ticker: "C T 1",
//			AllianceId: 1,
//		}
//
//		Convey("Corporations differ by name", func() {
//			esiCorporations := &chremoas_esi.Corporation{
//				Name: "Corporation Name 2",
//				Ticker: "C T 1",
//				AllianceId: 1,
//			}
//
//			answer := corporationDiffers(authCorporation, esiCorporations)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Corporations differ by ticker", func() {
//			esiCorporations := &chremoas_esi.Corporation{
//				Name: "Corporation Name 1",
//				Ticker: "C T 2",
//				AllianceId: 1,
//			}
//
//			answer := corporationDiffers(authCorporation, esiCorporations)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Corporations differ by alliance id", func() {
//			esiCorporations := &chremoas_esi.Corporation{
//				Name: "Corporation Name 1",
//				Ticker: "C T 1",
//				AllianceId: 2,
//			}
//
//			answer := corporationDiffers(authCorporation, esiCorporations)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Corporations are the same", func() {
//			esiCorporations := &chremoas_esi.Corporation{
//				Name: "Corporation Name 1",
//				Ticker: "C T 1",
//				AllianceId: 1,
//			}
//
//			answer := corporationDiffers(authCorporation, esiCorporations)
//
//			So(answer, ShouldEqual, false)
//		})
//	})
//
//	Convey("Should be able to compare 2 characters", t, func() {
//		authCharacter := &abaeve_auth.Character{
//			Id: 1,
//			Name: "Character Name 1",
//			CorporationId: 1,
//		}
//
//		Convey("Characters differ by name", func() {
//			esiCharacter := &chremoas_esi.Character{
//				Name: "Character Name 2",
//				CorporationId: 1,
//			}
//
//			answer := characterDiffers(authCharacter, esiCharacter)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Characters differ by corporation id", func() {
//			esiCharacter := &chremoas_esi.Character{
//				Name: "Character Name 1",
//				CorporationId: 2,
//			}
//
//			answer := characterDiffers(authCharacter, esiCharacter)
//
//			So(answer, ShouldEqual, true)
//		})
//
//		Convey("Characters are the same", func() {
//			esiCharacter := &chremoas_esi.Character{
//				Name: "Character Name 1",
//				CorporationId: 1,
//			}
//
//			answer := characterDiffers(authCharacter, esiCharacter)
//
//			So(answer, ShouldEqual, false)
//		})
//	})
//}
//
//func TestAuthEsiPoller_Poll(t *testing.T) {
//	SetDefaultFailureMode(FailureHalts)
//
//	Convey("Given the proper setup", t, func() {
//		mockEntityQueryClient, mockEntityAdminClient, mockAllianceClient, mockCorporationClient, mockCharacterClient := PollerTestSetup_Pegomock(t)
//		poller := NewAuthEsiPoller(mockEntityQueryClient, mockEntityAdminClient, mockAllianceClient, mockCorporationClient, mockCharacterClient)
//
//		Convey("Happy path", func() {
//			//<editor-fold desc="Common Given - mockAllianceClient.GetAllianceById, mockCorporationClient.GetCorporationById, mockCharacterClient.GetCharacterById stubs">
//			When(
//				mockAllianceClient.GetAllianceById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//						&chremoas_esi.GetAllianceByIdRequest{
//							Id: int32(1),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetAllianceByIdResponse{
//				Alliance: &chremoas_esi.Alliance{
//					Name: "Alliance Name 1",
//					Ticker: "A T 1",
//					DateFounded: int64(2),
//					ExecutorCorp: int32(2),
//				},
//			}, nil)
//
//			When(
//				mockCorporationClient.GetCorporationById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//						&chremoas_esi.GetCorporationByIdRequest{
//							Id: int32(1),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//				Corporation: &chremoas_esi.Corporation{
//					Name: "Corporation Name 1",
//					Ticker: "C T 1",
//					AllianceId: int32(1),
//					CeoId: int32(1),
//					CreationDate: int64(1),
//					CreatorId: int32(1),
//					Description: "Description for corp 1",
//					FactionId: int32(1),
//					MemberCount: int32(1),
//					TaxRate: float32(1),
//					Url: "Corp 1 url",
//				},
//			}, nil)
//
//			//Corporations
//			When(
//				mockCorporationClient.GetCorporationById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//						&chremoas_esi.GetCorporationByIdRequest{
//							Id: int32(2),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//				Corporation: &chremoas_esi.Corporation{
//					Name: "Corporation Name 2",
//					Ticker: "C T 2",
//					AllianceId: int32(1),
//					CeoId: int32(2),
//					CreationDate: int64(2),
//					CreatorId: int32(2),
//					Description: "Description for corp 2",
//					FactionId: int32(2),
//					MemberCount: int32(2),
//					TaxRate: float32(2),
//					Url: "Corp 2 url",
//				},
//			}, nil)
//
//			When(
//				mockCorporationClient.GetCorporationById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//						&chremoas_esi.GetCorporationByIdRequest{
//							Id: int32(3),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//				Corporation: &chremoas_esi.Corporation{
//					Name: "Corporation Name 3",
//					Ticker: "C T 3",
//					AllianceId: int32(2),
//					CeoId: int32(3),
//					CreationDate: int64(3),
//					CreatorId: int32(3),
//					Description: "Description for corp 3",
//					FactionId: int32(3),
//					MemberCount: int32(3),
//					TaxRate: float32(3),
//					Url: "Corp 3 url",
//				},
//			}, nil)
//
//			//Characters
//			When(
//				mockCharacterClient.GetCharacterById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//						&chremoas_esi.GetCharacterByIdRequest{
//							Id: int32(1),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//				Character: &chremoas_esi.Character{
//					Name: "Character Name 1",
//					Description: "Character 1 description",
//					CorporationId: int32(1),
//					AllianceId: int32(1),
//					AncestryId: int32(1),
//					Birthday: int64(1),
//					BloodlineId: int32(1),
//					Gender: "vOv",
//					RaceId: int32(1),
//					SecurityStatus: float32(-1),
//				},
//			}, nil)
//			When(
//				mockCharacterClient.GetCharacterById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//						&chremoas_esi.GetCharacterByIdRequest{
//							Id: int32(2),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//				Character: &chremoas_esi.Character{
//					Name: "Character Name 2",
//					Description: "Character 2 description",
//					CorporationId: int32(1),
//					AllianceId: int32(1),
//					AncestryId: int32(2),
//					Birthday: int64(2),
//					BloodlineId: int32(2),
//					Gender: "vOv",
//					RaceId: int32(2),
//					SecurityStatus: float32(-2),
//				},
//			}, nil)
//			When(
//				mockCharacterClient.GetCharacterById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//						&chremoas_esi.GetCharacterByIdRequest{
//							Id: int32(3),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//				Character: &chremoas_esi.Character{
//					Name: "Character Name 3",
//					Description: "Character 3 description",
//					CorporationId: int32(2),
//					AllianceId: int32(1),
//					AncestryId: int32(3),
//					Birthday: int64(3),
//					BloodlineId: int32(3),
//					Gender: "vOv",
//					RaceId: int32(3),
//					SecurityStatus: float32(-3),
//				},
//			}, nil)
//			When(
//				mockCharacterClient.GetCharacterById(
//					esi_matchers.AnyContextContext(),
//					esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//						&chremoas_esi.GetCharacterByIdRequest{
//							Id: int32(4),
//						},
//					),
//				),
//			).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//				Character: &chremoas_esi.Character{
//					Name: "Character Name 4",
//					Description: "Character 4 description",
//					CorporationId: int32(3),
//					AllianceId: int32(2),
//					AncestryId: int32(4),
//					Birthday: int64(4),
//					BloodlineId: int32(4),
//					Gender: "vOv",
//					RaceId: int32(4),
//					SecurityStatus: float32(-4),
//				},
//			}, nil)
//			//</editor-fold>
//
//			Convey("Alliance", func() {
//				//<editor-fold desc="Corporation and Character call ommitted from overall common responses">
//				When(
//					mockCorporationClient.GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(4),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//					Corporation: &chremoas_esi.Corporation{
//						Name: "Corporation Name 4",
//						Ticker: "C T 4",
//						CeoId: int32(4),
//						CreationDate: int64(4),
//						CreatorId: int32(4),
//						Description: "Description for corp 4",
//						FactionId: int32(4),
//						MemberCount: int32(4),
//						TaxRate: float32(4),
//						Url: "Corp 4 url",
//					},
//				},nil)
//
//				When(
//					mockCharacterClient.GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(5),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//					Character: &chremoas_esi.Character{
//						Name: "Character Name 5",
//						CorporationId: int32(4),
//						SecurityStatus: float32(-5),
//						RaceId: int32(5),
//						Gender: "male",
//						BloodlineId: int32(5),
//						Birthday: int64(5),
//						AncestryId: int32(5),
//						Description: "Character 5 description",
//					},
//				}, nil)
//				//</editor-fold>
//
//				Convey("ESI endpoints are called once for each known alliance and alliance 2 is removed", func() {
//					//<editor-fold desc="Test Specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockAllianceClient.GetAllianceById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//								&chremoas_esi.GetAllianceByIdRequest{
//									Id: int32(2),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetAllianceByIdResponse{}, nil)
//
//					When(
//						mockEntityAdminClient.AllianceUpdate(
//							authsrv_matchers.AnyContextContext(),
//							authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//								&abaeve_auth.AllianceAdminRequest{
//									Alliance: &abaeve_auth.Alliance{
//										Id:     int64(2),
//										Name:   "Alliance Name 2",
//										Ticker: "A T 2",
//									},
//									Operation: abaeve_auth.EntityOperation_REMOVE,
//								},
//							),
//						),
//					).ThenReturn(&abaeve_auth.EntityAdminResponse{
//						Success: true,
//					}, nil)
//					//</editor-fold>
//
//					//When
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then We expect only two calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockAllianceClient.VerifyWasCalledOnce().GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockAllianceClient.VerifyWasCalledOnce().GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//
//					//Enforce the 0 calls to Alliance 1 modifications
//					mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(1),
//									Name: "Alliance Name 1",
//									Ticker: "A T 1",
//								},
//								Operation: abaeve_auth.EntityOperation_REMOVE,
//							},
//						),
//					)
//					mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(1),
//									Name: "Alliance Name 1",
//									Ticker: "A T 1",
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//
//					//TODO: Pegomock doesn't handle varargs that aren't passed in without panicing
//					/*_, request, _ := */mockEntityAdminClient.VerifyWasCalledOnce().AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(2),
//									Name: "Alliance Name 2",
//									Ticker: "A T 2",
//								},
//								Operation: abaeve_auth.EntityOperation_REMOVE,
//							},
//						),
//					)/*.GetAllCapturedArguments()*/
//
//					/*
//					So(len(request), ShouldEqual, 1)
//					So(request[0], ShouldResemble,
//						&abaeve_auth.AllianceAdminRequest{
//							Alliance: &abaeve_auth.Alliance{
//								Id:     int64(2),
//								Name:   "Alliance Name 2",
//								Ticker: "A T 2",
//							},
//							Operation: abaeve_auth.EntityOperation_REMOVE,
//						},
//					)
//					*/
//					//</editor-fold>
//				})
//
//				Convey("ESI endpoints are called once for each known alliance and alliance 2 is renamed", func() {
//					//<editor-fold desc="Test specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockAllianceClient.GetAllianceById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//								&chremoas_esi.GetAllianceByIdRequest{
//									Id: int32(2),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetAllianceByIdResponse{
//						Alliance: &chremoas_esi.Alliance{
//							Name: "Alliance 2",
//							Ticker: "A T 2",
//							DateFounded: int64(2),
//							ExecutorCorp: int32(2),
//						},
//					}, nil)
//					//</editor-fold>
//
//					//When
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then We expect only two calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockAllianceClient.VerifyWasCalledOnce().GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockAllianceClient.VerifyWasCalledOnce().GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//
//					//Enforce the 0 calls to Alliance 1 modifications
//					mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(1),
//									Name: "Alliance Name 1",
//									Ticker: "A T 1",
//								},
//								Operation: abaeve_auth.EntityOperation_REMOVE,
//							},
//						),
//					)
//					mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(1),
//									Name: "Alliance Name 1",
//									Ticker: "A T 1",
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().AllianceUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//							&abaeve_auth.AllianceAdminRequest{
//								Alliance: &abaeve_auth.Alliance{
//									Id: int64(2),
//									Name: "Alliance 2",
//									Ticker: "A T 2",
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//			})
//
//			Convey("Corporation", func() {
//				//<editor-fold desc="Alliance and character call ommited from overall common responses">
//				When(
//					mockAllianceClient.GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(2),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetAllianceByIdResponse{
//					Alliance: &chremoas_esi.Alliance{
//						Name: "Alliance Name 2",
//						Ticker: "A T 2",
//						DateFounded: int64(2),
//						ExecutorCorp: int32(2),
//					},
//				}, nil)
//
//				When(
//					mockCharacterClient.GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(5),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//					Character: &chremoas_esi.Character{
//						Name: "Character Name 5",
//						CorporationId: int32(4),
//						SecurityStatus: float32(-5),
//						RaceId: int32(5),
//						Gender: "male",
//						BloodlineId: int32(5),
//						Birthday: int64(5),
//						AncestryId: int32(5),
//						Description: "Character 5 description",
//					},
//				}, nil)
//				//</editor-fold>
//
//				Convey("ESI endpoints are called for each known corporation and corporation 4 is removed", func() {
//					//<editor-fold desc="Test specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockCorporationClient.GetCorporationById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//								&chremoas_esi.GetCorporationByIdRequest{
//									Id: int32(4),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{},nil)
//					//</editor-fold>
//
//					//When
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 4 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CorporationUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//							&abaeve_auth.CorporationAdminRequest{
//								Corporation: &abaeve_auth.Corporation{
//									Id: int64(4),
//									Name: "Corporation Name 4",
//									Ticker: "C T 4",
//								},
//								Operation: abaeve_auth.EntityOperation_REMOVE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("ESI endpoints are called for each known corporation and corporation 4 is renamed", func() {
//					//<editor-fold desc="Test specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockCorporationClient.GetCorporationById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//								&chremoas_esi.GetCorporationByIdRequest{
//									Id: int32(4),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//						Corporation: &chremoas_esi.Corporation{
//							Name: "Corporation 4",
//							Ticker: "C T 4",
//							CeoId: int32(4),
//							CreationDate: int64(4),
//							CreatorId: int32(4),
//							Description: "Description for corp 4",
//							FactionId: int32(4),
//							MemberCount: int32(4),
//							TaxRate: float32(4),
//							Url: "Corp 4 url",
//						},
//					},nil)
//					//</editor-fold>
//
//					//When
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 4 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CorporationUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//							&abaeve_auth.CorporationAdminRequest{
//								Corporation: &abaeve_auth.Corporation{
//									Id: int64(4),
//									Name: "Corporation 4",
//									Ticker: "C T 4",
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("ESI endpoints are called for each known corporation and corporation 4 moves to alliance 1", func() {
//					//<editor-fold desc="Test specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockCorporationClient.GetCorporationById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//								&chremoas_esi.GetCorporationByIdRequest{
//									Id: int32(4),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//						Corporation: &chremoas_esi.Corporation{
//							Name: "Corporation Name 4",
//							Ticker: "C T 4",
//							AllianceId: int32(1),
//							CeoId: int32(4),
//							CreationDate: int64(4),
//							CreatorId: int32(4),
//							Description: "Description for corp 4",
//							FactionId: int32(4),
//							MemberCount: int32(4),
//							TaxRate: float32(4),
//							Url: "Corp 4 url",
//						},
//					},nil)
//					//</editor-fold>
//
//					//When
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 4 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CorporationUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//							&abaeve_auth.CorporationAdminRequest{
//								Corporation: &abaeve_auth.Corporation{
//									Id: int64(4),
//									Name: "Corporation Name 4",
//									AllianceId: int64(1),
//									Ticker: "C T 4",
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("Unknown alliance is created when a corporation is seen as joining it", func() {
//					//<editor-fold desc="Test specific givens">
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//					When(
//						mockAllianceClient.GetAllianceById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//								&chremoas_esi.GetAllianceByIdRequest{
//									Id: int32(3),
//								},
//							),
//						),
//					).ThenReturn(
//						&chremoas_esi.GetAllianceByIdResponse{
//							Alliance: &chremoas_esi.Alliance{
//								Name: "Alliance Name 3",
//								Ticker: "A T 3",
//								DateFounded: int64(3),
//								ExecutorCorp: int32(6),
//							},
//						},nil,
//					)
//
//					When(
//						mockCorporationClient.GetCorporationById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//								&chremoas_esi.GetCorporationByIdRequest{
//									Id: int32(4),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//						Corporation: &chremoas_esi.Corporation{
//							Name: "Corporation Name 4",
//							Ticker: "C T 4",
//							AllianceId: int32(3),
//							CeoId: int32(4),
//							CreationDate: int64(4),
//							CreatorId: int32(4),
//							Description: "Description for corp 4",
//							FactionId: int32(4),
//							MemberCount: int32(4),
//							TaxRate: float32(4),
//							Url: "Corp 4 url",
//						},
//					},nil)
//					//</editor-fold>
//
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 5 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockAllianceClient.VerifyWasCalledOnce().GetAllianceById(esi_matchers.AnyContextContext(), esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//						&chremoas_esi.GetAllianceByIdRequest{
//							Id: int32(3),
//						},
//					))
//
//					mockEntityAdminClient.VerifyWasCalledOnce().AllianceUpdate(authsrv_matchers.AnyContextContext(), authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//						&abaeve_auth.AllianceAdminRequest{
//							Alliance: &abaeve_auth.Alliance{
//								Id: int64(3),
//								Name: "Alliance Name 3",
//								Ticker: "A T 3",
//							},
//							Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//						},
//					))
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CorporationUpdate(authsrv_matchers.AnyContextContext(), authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//						&abaeve_auth.CorporationAdminRequest{
//							Corporation: &abaeve_auth.Corporation{
//								Id: int64(4),
//								Name: "Corporation Name 4",
//								Ticker: "C T 4",
//								AllianceId: int64(3),
//							},
//							Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//						},
//					))
//					//</editor-fold>
//				})
//			})
//
//			Convey("Character", func() {
//				//<editor-fold desc="Alliance and Corporation call ommited from overall common responses">
//				AuthSrvData_Pegomock(mockEntityQueryClient)
//				When(
//					mockAllianceClient.GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetAllianceByIdRequest(
//							&chremoas_esi.GetAllianceByIdRequest{
//								Id: int32(2),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetAllianceByIdResponse{
//					Alliance: &chremoas_esi.Alliance{
//						Name: "Alliance Name 2",
//						Ticker: "A T 2",
//						DateFounded: int64(2),
//						ExecutorCorp: int32(2),
//					},
//				}, nil)
//
//				When(
//					mockCorporationClient.GetCorporationById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//							&chremoas_esi.GetCorporationByIdRequest{
//								Id: int32(4),
//							},
//						),
//					),
//				).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//					Corporation: &chremoas_esi.Corporation{
//						Name: "Corporation Name 4",
//						Ticker: "C T 4",
//						CeoId: int32(4),
//						CreationDate: int64(4),
//						CreatorId: int32(4),
//						Description: "Description for corp 4",
//						FactionId: int32(4),
//						MemberCount: int32(4),
//						TaxRate: float32(4),
//						Url: "Corp 4 url",
//					},
//				},nil)
//				//</editor-fold>
//
//				Convey("ESI endpoints are called for each known character and character 5 is removed", func() {
//					//<editor-fold desc="Test specific givens">
//					When(
//						mockCharacterClient.GetCharacterById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//								&chremoas_esi.GetCharacterByIdRequest{
//									Id: int32(5),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{}, nil)
//					//</editor-fold>
//
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 5 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(5),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CharacterUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//							&abaeve_auth.CharacterAdminRequest{
//								Character: &abaeve_auth.Character{
//									Id: int64(5),
//									Name: "Character Name 5",
//									CorporationId: int64(4),
//								},
//								Operation: abaeve_auth.EntityOperation_REMOVE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("ESI endpoints are called for each known character and character 5 is renamed", func() {
//					//<editor-fold desc="Test specific givens">
//					When(
//						mockCharacterClient.GetCharacterById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//								&chremoas_esi.GetCharacterByIdRequest{
//									Id: int32(5),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//						Character: &chremoas_esi.Character{
//							Name: "Character 5",
//							CorporationId: int32(4),
//							SecurityStatus: float32(-5),
//							RaceId: int32(5),
//							Gender: "male",
//							BloodlineId: int32(5),
//							Birthday: int64(5),
//							AncestryId: int32(5),
//							Description: "Character 5 description",
//						},
//					}, nil)
//					//</editor-fold>
//
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 5 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(5),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CharacterUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//							&abaeve_auth.CharacterAdminRequest{
//								Character: &abaeve_auth.Character{
//									Id: int64(5),
//									Name: "Character 5",
//									CorporationId: int64(4),
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("ESI endpoints are called for each known character and character 5 moves to corporation 1", func() {
//					//<editor-fold desc="Test specific givens">
//					When(
//						mockCharacterClient.GetCharacterById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//								&chremoas_esi.GetCharacterByIdRequest{
//									Id: int32(5),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//						Character: &chremoas_esi.Character{
//							Name: "Character Name 5",
//							CorporationId: int32(1),
//							SecurityStatus: float32(-5),
//							RaceId: int32(5),
//							Gender: "male",
//							BloodlineId: int32(5),
//							Birthday: int64(5),
//							AncestryId: int32(5),
//							Description: "Character 5 description",
//						},
//					}, nil)
//					//</editor-fold>
//
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					//<editor-fold desc="Then we expect 5 calls to the esi endpoint and 1 call to the auth-srv endpoint">
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(1),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(2),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(3),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(4),
//							},
//						),
//					)
//					mockCharacterClient.VerifyWasCalledOnce().GetCharacterById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//							&chremoas_esi.GetCharacterByIdRequest{
//								Id: int32(5),
//							},
//						),
//					)
//
//					mockEntityAdminClient.VerifyWasCalledOnce().CharacterUpdate(
//						authsrv_matchers.AnyContextContext(),
//						authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//							&abaeve_auth.CharacterAdminRequest{
//								Character: &abaeve_auth.Character{
//									Id: int64(5),
//									Name: "Character Name 5",
//									CorporationId: int64(1),
//								},
//								Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//							},
//						),
//					)
//					//</editor-fold>
//				})
//
//				Convey("Unknown corporation is created when a character is seen as joining it", func() {
//					//<editor-fold desc="Test specific givens>
//					AuthSrvData_Pegomock(mockEntityQueryClient)
//
//					When(
//						mockCorporationClient.GetCorporationById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//								&chremoas_esi.GetCorporationByIdRequest{
//									Id: int32(5),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCorporationByIdResponse{
//						Corporation: &chremoas_esi.Corporation{
//							Name: "Corporation Name 5",
//							Ticker: "C T 5",
//							AllianceId: int32(2),
//							CeoId: int32(5),
//							CreationDate: int64(5),
//							CreatorId: int32(5),
//							Description: "Description for corp 5",
//							FactionId: int32(5),
//							MemberCount: int32(5),
//							TaxRate: float32(5),
//							Url: "Corp 5 url",
//						},
//					},nil)
//
//					When(
//						mockCharacterClient.GetCharacterById(
//							esi_matchers.AnyContextContext(),
//							esi_matchers.EqPtrToProtoGetCharacterByIdRequest(
//								&chremoas_esi.GetCharacterByIdRequest{
//									Id: int32(5),
//								},
//							),
//						),
//					).ThenReturn(&chremoas_esi.GetCharacterByIdResponse{
//						Character: &chremoas_esi.Character{
//							Name: "Character Name 5",
//							CorporationId: int32(5),
//							SecurityStatus: float32(-5),
//							RaceId: int32(5),
//							Gender: "male",
//							BloodlineId: int32(5),
//							Birthday: int64(5),
//							AncestryId: int32(5),
//							Description: "Character 5 description",
//						},
//					}, nil)
//					//</editor-fold>
//
//					err := poller.Poll()
//					So(err, ShouldBeNil)
//
//					mockCorporationClient.VerifyWasCalledOnce().GetCorporationById(esi_matchers.AnyContextContext(), esi_matchers.EqPtrToProtoGetCorporationByIdRequest(
//						&chremoas_esi.GetCorporationByIdRequest{
//							Id: int32(5),
//						},
//					))
//					mockEntityAdminClient.VerifyWasCalledOnce().CorporationUpdate(authsrv_matchers.AnyContextContext(), authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//						&abaeve_auth.CorporationAdminRequest{
//							Corporation: &abaeve_auth.Corporation{
//								Id: 5,
//								Name: "Corporation Name 5",
//								Ticker: "C T 5",
//								AllianceId: 2,
//							},
//							Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//						},
//					))
//					mockEntityAdminClient.VerifyWasCalledOnce().CharacterUpdate(authsrv_matchers.AnyContextContext(), authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//						&abaeve_auth.CharacterAdminRequest{
//							Character: &abaeve_auth.Character{
//								Id: 5,
//								Name: "Character Name 5",
//								CorporationId: 5,
//							},
//							Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//						},
//					))
//				})
//
//				Convey("Unknown alliance of unknown corporation is created when a character is seen as joining it", func() {
//
//				})
//			})
//		})
//
//		//TODO: Complete these
//		Convey("Error paths", func() {
//			Convey("ESI", func() {
//				When(
//					mockAllianceClient.GetAllianceById(
//						esi_matchers.AnyContextContext(),
//						esi_matchers.AnyPtrToProtoGetAllianceByIdRequest(),
//					),
//				).ThenReturn(nil, errors.New("bob, I failed you on the ESI calls :("))
//
//				Convey("Alliance", func() {
//					Convey("ESI endpoints return an error and no alliances are udpated/removed", func() {
//
//					})
//				})
//
//				Convey("Corporation", func() {
//
//				})
//
//				Convey ("Character", func() {
//
//				})
//			})
//
//			Convey("Auth Srv", func() {
//				Convey("Alliance", func() {
//					Convey("ESI endpoints called once for each known alliance and alliance update/remove returns error", func() {
//
//					})
//				})
//
//				Convey("Corporation", func() {
//
//				})
//
//				Convey ("Character", func() {
//
//				})
//			})
//		})
//
//		//<editor-fold desc="Enforce alliance 1 is never updated or deleted, corporations 1-3 are never updated or deleted and characters 1-4 are never updated or deleted">
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//				&abaeve_auth.AllianceAdminRequest{
//					Alliance: &abaeve_auth.Alliance{
//						Id: int64(1),
//						Name: "Alliance 1",
//						Ticker: "A T 1",
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).AllianceUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoAllianceAdminRequest(
//				&abaeve_auth.AllianceAdminRequest{
//					Alliance: &abaeve_auth.Alliance{
//						Id: int64(1),
//						Name: "Alliance Name 1",
//						Ticker: "A T 1",
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(1),
//						Name: "Corporation 1",
//						Ticker: "C T 1",
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(2),
//						Name: "Corporation 2",
//						Ticker: "C T 2",
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(3),
//						Name: "Corporation 3",
//						Ticker: "C T 3",
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(1),
//						Name: "Corporation Name 1",
//						Ticker: "C T 1",
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(2),
//						Name: "Corporation Name 2",
//						Ticker: "C T 2",
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CorporationUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCorporationAdminRequest(
//				&abaeve_auth.CorporationAdminRequest{
//					Corporation: &abaeve_auth.Corporation{
//						Id: int64(3),
//						Name: "Corporation Name 3",
//						Ticker: "C T 3",
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(1),
//						Name: "Character 1",
//						CorporationId: int64(1),
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(2),
//						Name: "Character 2",
//						CorporationId: int64(2),
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(3),
//						Name: "Character 3",
//						CorporationId: int64(2),
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(4),
//						Name: "Character 4",
//						CorporationId: int64(3),
//					},
//					Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
//				},
//			),
//		)
//
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(1),
//						Name: "Character Name 1",
//						CorporationId: int64(1),
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(2),
//						Name: "Character Name 2",
//						CorporationId: int64(2),
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(3),
//						Name: "Character Name 3",
//						CorporationId: int64(2),
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		mockEntityAdminClient.VerifyWasCalled(AtMost(0)).CharacterUpdate(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoCharacterAdminRequest(
//				&abaeve_auth.CharacterAdminRequest{
//					Character: &abaeve_auth.Character{
//						Id: int64(4),
//						Name: "Character Name 1",
//						CorporationId: int64(3),
//					},
//					Operation: abaeve_auth.EntityOperation_REMOVE,
//				},
//			),
//		)
//		//</editor-fold>
//	})
//}
//
////Our Mock auth-srv is _always_ going to have the same set of data for all of our tests so it's expectations and responses will always be the same
//// This data is as follows:
//// Alliances:
////		One:
////			ID: 1
////			Name: Alliance Name 1
////			Ticker: A T 1
////		Two:
////			ID: 2
////			Name: Alliance Name 2
////			Ticker: A T 2
////
//// Corporations:
////		One:
////			ID: 1
////			Name: Corporation Name 1
////			Ticker: C T 1
////			Alliance ID: 1
////		Two:
////			ID: 2
////			Name: Corporation Name 2
////			Ticker: C T 2
////			Alliance ID: 1
////		Three:
////			ID: 3
////			Name: Corporation Name 3
////			Ticker: C T 3
////			Alliance ID: 2
////		Four:
////			ID: 4
////			Name: Corporation Name 4
////			Ticker: C T 4
////			Alliance ID: N/A
////
//// Characters:
////		One:
////			ID: 1
////			Name: Character Name 1
////			Corporation ID: 1
////		Two:
////			ID: 2
////			Name: Character Name 2
////			Corporation ID: 1
////		Three:
////			ID: 3
////			Name: Character Name 3
////			Corporation ID: 2
////		Four:
////			ID: 4
////			Name: Character Name 4
////			Corporation ID: 3
////		Five:
////			ID: 5
////			Name: Character Name 5
////			Corporation ID: 4
//func AuthSrvData_Pegomock(mockEntityQueryClient *authsrv_mocks.MockEntityQueryClient) {
//	When(
//		mockEntityQueryClient.GetAlliances(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoEntityQueryRequest(
//				&abaeve_auth.EntityQueryRequest{
//					EntityType: abaeve_auth.EntityType_ALLIANCE,
//				},
//			),
//		),
//	).ThenReturn(&abaeve_auth.AlliancesResponse{
//		List: []*abaeve_auth.Alliance{
//			{
//				Id:     int64(1),
//				Name:   "Alliance Name 1",
//				Ticker: "A T 1",
//			},
//			{
//				Id:     int64(2),
//				Name:   "Alliance Name 2",
//				Ticker: "A T 2",
//			},
//		},
//	}, nil)
//
//	When(
//		mockEntityQueryClient.GetCorporations(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoEntityQueryRequest(
//				&abaeve_auth.EntityQueryRequest{
//					EntityType: abaeve_auth.EntityType_CORPORATION,
//				},
//			),
//		),
//	).ThenReturn(
//		&abaeve_auth.CorporationsResponse{
//			List: []*abaeve_auth.Corporation{
//				{
//					Id: int64(1),
//					Name: "Corporation Name 1",
//					Ticker: "C T 1",
//					AllianceId: int64(1),
//				},
//				{
//					Id: int64(2),
//					Name: "Corporation Name 2",
//					Ticker: "C T 2",
//					AllianceId: int64(1),
//				},
//				{
//					Id: int64(3),
//					Name: "Corporation Name 3",
//					Ticker: "C T 3",
//					AllianceId: int64(2),
//				},
//				{
//					Id: int64(4),
//					Name: "Corporation Name 4",
//					Ticker: "C T 4",
//				},
//			},
//		}, nil,
//	)
//
//	When(
//		mockEntityQueryClient.GetCharacters(
//			authsrv_matchers.AnyContextContext(),
//			authsrv_matchers.EqPtrToProtoEntityQueryRequest(
//				&abaeve_auth.EntityQueryRequest{
//					EntityType: abaeve_auth.EntityType_CHARACTER,
//				},
//			),
//		),
//	).ThenReturn(
//		&abaeve_auth.CharactersResponse{
//			List: []*abaeve_auth.Character{
//				{
//					Id: int64(1),
//					Name: "Character Name 1",
//					CorporationId: 1,
//				},
//				{
//					Id: int64(2),
//					Name: "Character Name 2",
//					CorporationId: 1,
//				},
//				{
//					Id: int64(3),
//					Name: "Character Name 3",
//					CorporationId: 2,
//				},
//				{
//					Id: int64(4),
//					Name: "Character Name 4",
//					CorporationId: 3,
//				},
//				{
//					Id: int64(5),
//					Name: "Character Name 5",
//					CorporationId: 4,
//				},
//			},
//		}, nil,
//	)
//}
//
//func NewFailHandler(t *testing.T) FailHandler {
//	return func(message string, callerSkip ...int) {
//		skip := 1
//		if len(callerSkip) > 0 {
//			skip = callerSkip[0]
//		}
//		stackTrace := pruneStack(string(debug.Stack()), skip)
//		t.Errorf("\n%s\n\n%s\n", message, stackTrace)
//	}
//}
//
//func pruneStack(fullStackTrace string, skip int) string {
//	stack := strings.Split(fullStackTrace, "\n")
//	if len(stack) > skip+4 {
//		stack = stack[skip+4:]
//	}
//	prunedStack := []string{}
//	re := regexp.MustCompile(`/auth-esi-poller/`)
//	for i := 0; i < len(stack); i += 2 {
//		if re.Match([]byte(stack[i])) {
//			prunedStack = append(prunedStack, stack[i+1])
//			break
//		}
//	}
//	return strings.Join(prunedStack, "\n")
//}
