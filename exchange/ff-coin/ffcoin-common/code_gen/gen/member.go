package gen


type Member struct {
    Id  int64  `json:"id" from:"id"`
    AliNo  string  `json:"aliNo" from:"aliNo"`
    QrCodeUrl  string  `json:"qrCodeUrl" from:"qrCodeUrl"`
    AppealSuccessTimes  int  `json:"appealSuccessTimes" from:"appealSuccessTimes"`
    AppealTimes  int  `json:"appealTimes" from:"appealTimes"`
    ApplicationTime  int64  `json:"applicationTime" from:"applicationTime"`
    Avatar  string  `json:"avatar" from:"avatar"`
    Bank  string  `json:"bank" from:"bank"`
    Branch  string  `json:"branch" from:"branch"`
    CardNo  string  `json:"cardNo" from:"cardNo"`
    CertifiedBusinessApplyTime  int64  `json:"certifiedBusinessApplyTime" from:"certifiedBusinessApplyTime"`
    CertifiedBusinessCheckTime  int64  `json:"certifiedBusinessCheckTime" from:"certifiedBusinessCheckTime"`
    CertifiedBusinessStatus  int  `json:"certifiedBusinessStatus" from:"certifiedBusinessStatus"`
    ChannelId  int  `json:"channelId" from:"channelId"`
    Email  string  `json:"email" from:"email"`
    FirstLevel  int  `json:"firstLevel" from:"firstLevel"`
    GoogleDate  int64  `json:"googleDate" from:"googleDate"`
    GoogleKey  string  `json:"googleKey" from:"googleKey"`
    GoogleState  int  `json:"googleState" from:"googleState"`
    IdNumber  string  `json:"idNumber" from:"idNumber"`
    InviterId  int64  `json:"inviterId" from:"inviterId"`
    IsChannel  int  `json:"isChannel" from:"isChannel"`
    JyPassword  string  `json:"jyPassword" from:"jyPassword"`
    LastLoginTime  int64  `json:"lastLoginTime" from:"lastLoginTime"`
    City  string  `json:"city" from:"city"`
    Country  string  `json:"country" from:"country"`
    District  string  `json:"district" from:"district"`
    Province  string  `json:"province" from:"province"`
    LoginCount  int  `json:"loginCount" from:"loginCount"`
    LoginLock  int  `json:"loginLock" from:"loginLock"`
    Margin  string  `json:"margin" from:"margin"`
    MemberLevel  int  `json:"memberLevel" from:"memberLevel"`
    MobilePhone  string  `json:"mobilePhone" from:"mobilePhone"`
    Password  string  `json:"password" from:"password"`
    PromotionCode  string  `json:"promotionCode" from:"promotionCode"`
    PublishAdvertise  int  `json:"publishAdvertise" from:"publishAdvertise"`
    RealName  string  `json:"realName" from:"realName"`
    RealNameStatus  int  `json:"realNameStatus" from:"realNameStatus"`
    RegistrationTime  int64  `json:"registrationTime" from:"registrationTime"`
    Salt  string  `json:"salt" from:"salt"`
    SecondLevel  int  `json:"secondLevel" from:"secondLevel"`
    SignInAbility  int  `json:"signInAbility" from:"signInAbility"`
    Status  int  `json:"status" from:"status"`
    ThirdLevel  int  `json:"thirdLevel" from:"thirdLevel"`
    Token  string  `json:"token" from:"token"`
    TokenExpireTime  int64  `json:"tokenExpireTime" from:"tokenExpireTime"`
    TransactionStatus  int  `json:"transactionStatus" from:"transactionStatus"`
    TransactionTime  int64  `json:"transactionTime" from:"transactionTime"`
    Transactions  int  `json:"transactions" from:"transactions"`
    Username  string  `json:"username" from:"username"`
    QrWeCodeUrl  string  `json:"qrWeCodeUrl" from:"qrWeCodeUrl"`
    Wechat  string  `json:"wechat" from:"wechat"`
    Local  string  `json:"local" from:"local"`
    Integration  int64  `json:"integration" from:"integration"`
    MemberGradeId  int64  `json:"memberGradeId" from:"memberGradeId"`
    KycStatus  int  `json:"kycStatus" from:"kycStatus"`
    GeneralizeTotal  int64  `json:"generalizeTotal" from:"generalizeTotal"`
    InviterParentId  int64  `json:"inviterParentId" from:"inviterParentId"`
    SuperPartner  string  `json:"superPartner" from:"superPartner"`
    KickFee  float64  `json:"kickFee" from:"kickFee"`
    Power  float64  `json:"power" from:"power"`
    TeamLevel  int  `json:"teamLevel" from:"teamLevel"`
    TeamPower  float64  `json:"teamPower" from:"teamPower"`
    MemberLevelId  int64  `json:"memberLevelId" from:"memberLevelId"`
}