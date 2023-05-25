package clothing

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/credential"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/auth"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/subscribe"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbatch"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"strings"
	"time"
)

type WechatService struct {
}

func (w *WechatService) RefreshAccessToken(mini *miniprogram.MiniProgram) error {
	defaultAccessToken := mini.GetContext().AccessTokenHandle.(*credential.DefaultAccessToken)
	err := defaultAccessToken.RefreshAccessToken()
	if err != nil {
		return err
	}
	return nil
}

func (w *WechatService) GetWechatSession(jsCode string, appID int) (session auth.ResCode2Session, err error) {

	a := global.GVA_WECHAT_MINI.GetAuth()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	session, err = a.Code2SessionContext(ctx, jsCode)
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
	}
	return
}

func (w *WechatService) GetWechatPhoneNumber(jsCode string, appID int) (phone *auth.GetPhoneNumberResponse, err error) {
	if jsCode == "" {
		return nil, errors.New("用户拒绝授权")
	}
	mini := global.GVA_WECHAT_MINI
	a := global.GVA_WECHAT_MINI.GetAuth()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	phone, err = a.GetPhoneNumberContext(ctx, jsCode)
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		if err := w.RefreshAccessToken(mini); err == nil {
			phone, err = a.GetPhoneNumberContext(ctx, jsCode)
			if err != nil {
				global.GVA_LOG.Sugar().Error(err)
				return nil, errors.New("获取失败请重试")
			}
		} else {
			global.GVA_LOG.Sugar().Error(err)
		}
	}
	return
}

func (w *WechatService) SendDrawSubscribeMsg(openID, msg string) {
	// 获取小程序client发送订阅消息
	mini := global.GVA_WECHAT_MINI
	var miniMessage subscribe.Message
	miniMessage.ToUser = openID
	miniMessage.TemplateID = "E_CWXNUjrdzLA3Xxd3Y6SJ9nV1tu9gRzJ7clW2497xk"
	miniMessage.Data = map[string]*subscribe.DataItem{
		"thing7": {
			Value: "AI画画",
			Color: "",
		},
		"date2": {
			Value: time.Now().Format("2006-01-02 15:04:05"),
			Color: "",
		},
		"thing3": {
			Value: msg,
			Color: "",
		},
	}
	miniMessage.Page = "/pages/drawList/drawList"
	miniMessage.MiniprogramState = global.GVA_CONFIG.Wechat.MiniprogramState
	miniMessage.Lang = "zh_CN"
	if err := mini.GetSubscribe().Send(&miniMessage); err != nil {
		global.GVA_LOG.Sugar().Error("Send Subscribe fail ", err.Error())
		if strings.Contains(err.Error(), "40001") {
			if err := w.RefreshAccessToken(mini); err == nil {
				if err := mini.GetSubscribe().Send(&miniMessage); err != nil {
					global.GVA_LOG.Sugar().Error("Send Subscribe fail again ", err.Error())
				}
			} else {
				global.GVA_LOG.Sugar().Error(err)
			}
		}
	}
}

func (w *WechatService) CreateJSAPIOrder(openID, appID, Description, Attach string, order clothing.Order) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	//if len(openID) == 0 {
	//	var openIDObj clothing.UserOpenID
	//	global.GVA_DB.Where("user_id = ? and app = ?", order.UserID, appID).First(&openIDObj)
	//	openID = openIDObj.OpenID
	//}
	resp, _, err := global.GVA_WECHAT_PAY.PrepayWithRequestPayment(context.Background(),
		jsapi.PrepayRequest{
			Appid:       core.String(appID),
			Mchid:       core.String(global.GVA_CONFIG.Wechat.MchID),
			Description: core.String(Description),
			OutTradeNo:  core.String(order.OrderNo),
			Attach:      core.String(Attach),
			NotifyUrl:   core.String(global.GVA_CONFIG.Wechat.PayNotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(int64(order.Amount * 100)),
				//Total: core.Int64(1),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openID),
			},
		},
	)
	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

func (w *WechatService) CreateAppOrder(appID, Description, Attach string, order clothing.Order) (*app.PrepayWithRequestPaymentResponse, error) {
	resp, _, err := global.GVA_WECHAT_PAY_APP.PrepayWithRequestPayment(context.Background(), app.PrepayRequest{
		Appid:       core.String(appID),
		Mchid:       core.String(global.GVA_CONFIG.Wechat.MchID),
		Description: core.String(Description),
		OutTradeNo:  core.String(order.OrderNo),
		Attach:      core.String(Attach),
		NotifyUrl:   core.String(global.GVA_CONFIG.Wechat.PayNotifyUrl),
		Amount: &app.Amount{
			Total: core.Int64(int64(order.Amount * 100)),
		},
	})
	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

func (w *WechatService) CreateH5Order(appID, Description, Attach string, order clothing.Order) (*h5.PrepayResponse, error) {
	resp, _, err := global.GVA_WECHAT_PAY_H5.Prepay(context.Background(), h5.PrepayRequest{
		Appid:       core.String(appID),
		Mchid:       core.String(global.GVA_CONFIG.Wechat.MchID),
		Description: core.String(Description),
		OutTradeNo:  core.String(order.OrderNo),
		Attach:      core.String(Attach),
		NotifyUrl:   core.String(global.GVA_CONFIG.Wechat.PayNotifyUrl),
		Amount: &h5.Amount{
			Total: core.Int64(int64(order.Amount * 100)),
		},
	})
	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

func (w *WechatService) WxPayCallback(c *gin.Context) (*payments.Transaction, error) {
	transaction := new(payments.Transaction)
	request, err := global.GVA_WECHAT_PAY_HANDLER.ParseNotifyRequest(context.Background(), c.Request, transaction)
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return nil, err
	}
	// 处理通知内容
	global.GVA_LOG.Sugar().Info(request.Summary)
	return transaction, err
}

type Refund struct {
	Mchid               string     `json:"mchid"`
	TransactionId       string     `json:"transaction_id"`
	OutTradeNo          string     `json:"out_trade_no"`
	RefundId            string     `json:"refund_id"`
	OutRefundNo         string     `json:"out_refund_no"`
	RefundStatus        string     `json:"refund_status"`
	SuccessTime         *time.Time `json:"success_time"`
	UserReceivedAccount string     `json:"user_received_account"`
	Amount              struct {
		Total       int `json:"total"`
		Refund      int `json:"refund"`
		PayerTotal  int `json:"payer_total"`
		PayerRefund int `json:"payer_refund"`
	} `json:"amount"`
}

func (w *WechatService) WxPayRefundCallback(c *gin.Context) (*Refund, error) {
	refund := new(Refund)
	request, err := global.GVA_WECHAT_PAY_HANDLER.ParseNotifyRequest(context.Background(), c.Request, refund)
	if err != nil {
		global.GVA_LOG.Sugar().Error(err)
		return nil, err
	}
	// 处理通知内容
	global.GVA_LOG.Sugar().Info(request.Summary)
	return refund, err
}

func (w *WechatService) WxSendRedPack() {
	var (
		mchID                      = "190000****"                               // 商户号
		mchCertificateSerialNumber = "3775************************************" // 商户证书序列号
		mchAPIv3Key                = "2ab9****************************"         // 商户APIv3密钥
	)
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}
	svc := transferbatch.TransferBatchApiService{Client: client}
	resp, result, err := svc.InitiateBatchTransfer(ctx,
		transferbatch.InitiateBatchTransferRequest{
			Appid:       core.String("wxf636efh567hg4356"),
			OutBatchNo:  core.String("plfk2020042013"),
			BatchName:   core.String("2019年1月深圳分部报销单"),
			BatchRemark: core.String("2019年1月深圳分部报销单"),
			TotalAmount: core.Int64(4000000),
			TotalNum:    core.Int64(200),
			TransferDetailList: []transferbatch.TransferDetailInput{transferbatch.TransferDetailInput{
				OutDetailNo:    core.String("x23zy545Bd5436"),
				TransferAmount: core.Int64(200000),
				TransferRemark: core.String("2020年4月报销"),
				Openid:         core.String("o-MYE42l80oelYMDE34nYD456Xoy"),
				UserName:       core.String("757b340b45ebef5467rter35gf464344v3542sdf4t6re4tb4f54ty45t4yyry45"),
				UserIdCard:     core.String("8609cb22e1774a50a930e414cc71eca06121bcd266335cda230d24a7886a8d9f"),
			}},
		},
	)
	if err != nil {
		// 处理错误
		log.Printf("call InitiateBatchTransfer err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
