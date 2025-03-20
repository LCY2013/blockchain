<template>
  <div class="activity">
    <img style="banner" src="../../assets/images/crowdfunding_banner.png">
    <div class="banner_text">
      <p>{{$t('crowdfunding.totalAmount')}}</p>
      <p>
        <span class="money">{{totalMoney}}</span>
        <br class="br">
        <span class="money_span">{{$t('crowdfunding.totalAmount2')}}</span>
      </p>
      <div class="banner_content">
        <p>{{totalTimes}}<br class="br"><span>{{$t('crowdfunding.donations')}}</span></p>
        <p>{{totalProject}}<br class="br"><span>{{$t('crowdfunding.projectsGetHelp')}}</span></p>
      </div>
    </div>
    <div class="tab">
      <Tabs value="name1" class="tabs1">
        <TabPane :label="$t('crowdfunding.crowdfunding')" name="name1">
          <div class="tab_all" @click="syadd_show()">{{$t('crowdfunding.seeMore')}} ></div>
          <div class="swiper-container">  
            <div class="swiper-wrapper">
              <div class="swiper-slide swiper-slide1" v-for="(item, index) in list.slice(0, 4)" :key="index">
                <div @click="xs_details_show(item)">
                  <img :src="item.picUrl"/>
                  <div class="swiper_text">
                    <div class="swiper_title">
                      <p>{{item.fundingTitle}}</p>
                      <p>{{dateFormat(item.addTime)}}—{{dateFormat(item.endTime)}}</p>
                    </div>
                    <Row class="swiper_content">
                        <Col span="8">
                          <p>{{item.targetAmount}}</p>
                          <p>{{$t('crowdfunding.targetAmount')}}(USDT)</p>
                        </Col>
                        <Col span="8">
                          <p>{{item.amountReceived}}</p>
                          <p>{{$t('crowdfunding.amountRaised')}}(USDT)</p>
                        </Col>
                        <Col span="8">
                          <p>{{item.times}}</p>
                          <p>{{$t('crowdfunding.helpTimes')}}({{$t('crowdfunding.times')}})</p>
                        </Col>
                    </Row>
                  </div>
                </div>
              </div>
            </div>
            <div class="swiper-pagination"></div>
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
          </div>
          <div class="tab_all2" @click="syadd_show()">{{$t('crowdfunding.seeMore')}}</div>
          <div class="btn" v-if="ifxsBtn != ''" @click="maskClick()">
              <span>{{$t('crowdfunding.initiateCrowdfunding')}}</span>
          </div>
        </TabPane>
        <TabPane :label="$t('crowdfunding.publicWelfare')" name="name2">
          <div class="tab_all" @click="syadd_show2()">{{$t('crowdfunding.seeMore')}} ></div>
          <div class="swiper-container">
            <div class="swiper-wrapper">
              <div class="swiper-slide swiper-slide2" v-for="(item, index) in enablewelfare.slice(0, 4)" :key="index">
                <div @click="xx_detailsClick(item)">
                  <img :src="item.picUrl"/>
                  <div class="swiper_text">
                    <div class="swiper_title">
                      <p>{{item.fundingTitle}}</p>
                      <p>{{dateFormat(item.passTime)}}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="swiper-pagination"></div>
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
          </div>
          <div class="tab_all2" @click="syadd_show2()">{{$t('crowdfunding.seeMore')}}</div>
          <div class="btn" v-if="ifxxBtn != ''" @click="xx_apply()">
              <span>{{$t('crowdfunding.applySubsidy')}}</span>
          </div>
        </TabPane>
      </Tabs>
      <div class="bottom">
        <img src="../../assets/images/crowdfunding_text.png"/>
      </div>
      <div class="bottom_text">
        <p>{{$t('crowdfunding.loveReminder')}}</p>
        <p>{{$t('crowdfunding.loveReminder2')}}</p>
      </div>
    </div>
    <!--线上全部-->
    <div class="Initiate_mask4" v-if="all_show">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.allCrowdfunding')}}</p>
          <img @click="allClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content" style="max-height: 570px; overflow-y: scroll;">
          <div class="all_xs"  @click="xs_details_show2(item)" v-for="(item, index) in list" :key="index">
            <div class="all_xs_img">
              <img :src="item.picUrl"/>
            </div>
            <div class="all_xs_text">
              <div class="all_xs_text_left">
                <p class="all_xs_text_title">{{item.fundingTitle}}</p>
                <div class="all_xs_text_money">
                  <div class="all_xs_text_content">
                        <div>
                          <p>{{$t('crowdfunding.amountRaised')}}<span>(USDT)</span></p>
                          <p>{{item.amountReceived}}</p>
                        </div>
                        <div>
                          <p class="all_xs_text_money">======</p>
                        </div>
                        <div>
                          <p>{{$t('crowdfunding.targetAmount')}}<span>(USDT)</span></p>
                          <p>{{item.targetAmount}}</p>
                        </div>
                    </div>
                </div>
              </div>
              <div class="all_xs_text_right">
                <div class="all_xs_right_top">
                  <span>{{item.times}}</span>
                  <span>{{$t('crowdfunding.helpTimes')}}</span>
                </div>
                <div class="all_xs_right_bottom">
                  {{dateFormat(item.passTime)}}—{{dateFormat(item.endTime)}}
                </div>
              </div>
              <div style="clear:both;"></div>
            </div>
            <div style="clear:both;"></div>
          </div>
        </div>
      </div>
    </div>
    <!--线下全部-->
    <div class="Initiate_mask4" v-if="all_show2">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.allPublicWelfare')}}</p>
          <img @click="allClick2()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content" style="max-height: 570px; overflow-y: scroll;" >
          <div class="all_xs" @click="xx_detailsClick2(item)" v-for="(item, index) in enablewelfare" :key="index">
            <div class="all_xs_img">
              <img :src="item.picUrl"/>
            </div>
            <div class="all_xs_text">
              <div class="all_xs_text_left">
                <p class="all_xs_text_title">{{item.fundingTitle}}</p>
                <div class="all_xs_text_money">
                  <div class="all_xs_text_content">
                        <div>
                          <p>{{$t('crowdfunding.auditAmount')}}<span>(USDT)</span></p>
                          <p>{{item.amountReceived}}</p>
                        </div>
                        <div>
                          <p class="all_xs_text_money">======</p>
                        </div>
                        <div>
                          <p>{{$t('crowdfunding.applicationAmount')}}<span>(USDT)</span></p>
                          <p>{{item.applyAmount}}</p>
                        </div>
                    </div>
                </div>
              </div>
              <div class="all_xs_text_right">
                <div class="all_xs_right_top">
                  {{dateFormat(item.passTime)}}
                </div>
              </div>
              <div style="clear:both;"></div>
            </div>
            <div style="clear:both;"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="mask" v-if="mask"></div>
    <!--选择捐款类型-->
    <div class="Initiate_mask" v-if="Initiate_select">
      <div class="Initiate_crowdfunding">
        <div class="select_title">
          <p>{{$t('crowdfunding.fundraisingType')}}</p>
          <img @click="maskClick2()" src="../../assets/images/down.png" />
        </div>
        <Row :gutter="10">
            <i-col :xs="{ span: 8 }" :lg="{ span: 6}">
                <Card shadow>
                  <div class="window" @click="selectClick1()">
                    <img src="../../assets/images/crowdfunding_tk1.png" />
                    <p>{{$t('crowdfunding.wish')}}</p>
                  </div>
                </Card>
            </i-col>
            <i-col :xs="{ span: 8 }" :lg="{ span: 6, offset: 3 }">
                <Card shadow>
                  <div class="window" @click="selectClick2()">
                    <img class="window_img_yl" src="../../assets/images/crowdfunding_tk2.png" />
                    <p>{{$t('crowdfunding.medical')}}</p>
                  </div>
                </Card>
            </i-col>
            <i-col :xs="{ span: 8 }" :lg="{ span: 6, offset: 3 }">
                <Card shadow>
                  <div class="window" @click="selectClick3()">
                    <img src="../../assets/images/crowdfunding_tk3.png" />
                    <p>{{$t('crowdfunding.pioneer')}}</p>
                  </div>
                </Card>
            </i-col>
        </Row>
      </div>
    </div>
    <!--在线心愿-->
    <div class="Initiate_mask2" v-if="Initiate_xy">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.basicInformation')}}</p>
          <img @click="showClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_tk_nr">
            <div style="width:350px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.demandAmount')}}</p>
              <i-input v-model="xyDetails.targetAmount" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
                  <span slot="append" style="color:red;margin-left:10px;">{{$t('crowdfunding.maxMoney')}} {{level}} {{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.fundraisingTitle')}}</p>
              <i-input v-model="xyDetails.fundingTitle" placeholder="一句话简单描述你的情况(25个字以内）"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.instructions')}}</p>
              <i-input v-model="xyDetails.instructions" type="textarea" :rows="4" placeholder="可为自己/家人/其他人发起该心愿众筹活动，
              请在此栏填写详细的自我介绍/再将心愿梦想完整准确的表述出来以获得金融家们的支持。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />  {{$t('crowdfunding.addImg')}}</p>
              <Upload ref="upload1" :on-success="zcImage" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
              <!-- <div class="img_text">建议您上传患者积极乐观向上的医院治疗照</div> -->
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.sponsorInfo')}}</p>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorName')}}</p>
                <i-input v-model="xyDetails.sponsor" placeholder="请输入发起人的姓名"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorType')}}</p>
                <RadioGroup v-model="xyDetails.documentType">
                  <Radio label="1" value="1">{{$t('crowdfunding.idCard')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.passport')}}</Radio>
                </RadioGroup>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.idNumber')}}</p>
                <i-input v-model="xyDetails.documentNumber" placeholder="请输入证件号码"></i-input>
              </div>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.otherMaterials')}}</p>
              <Upload ref="upload2" :multiple="true" :on-success="zcImage2" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
        </div>
        <div class="window_btn" @click="xyClick(type = 1)">{{$t('crowdfunding.submit')}}</div>
      </div>
    </div>
    <!--在线医疗-->
    <div class="Initiate_mask2" v-if="Initiate_yl">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.basicInformation')}}</p>
          <img @click="showClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_tk_nr">
            <div style="width:350px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.demandAmount')}}</p>
              <i-input placeholder="请输入金额" v-model="ylDetails.targetAmount" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
                  <span slot="append" style="color:red;margin-left:10px;">{{$t('crowdfunding.maxMoney')}} {{level}} {{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.fundraisingTitle')}}</p>
              <i-input  v-model="ylDetails.fundingTitle" placeholder="一句话简单描述你的情况(25个字以内）"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.fundraisingUse')}}</p>
              <i-input v-model="ylDetails.fundingUse" type="textarea" placeholder="可为自己/家人/其他人发起该医疗众筹活动，请在此栏详细填写你筹集该笔资金的各种用途。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.plan')}}</p>
              <i-input v-model="ylDetails.plan" type="textarea" placeholder="请在此栏详细填写该笔资金受益人在获得资金后的执行计划。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.projectStory')}}</p>
              <i-input v-model="ylDetails.story" type="textarea" :rows="4" placeholder="我叫xxx，今年xxx岁，xx省xx市人。职业是xxx。xx年xx月xx日，在xx医院，被确诊为xx疾病。从确诊为xx疾病。从确
诊到现在已经花费了xx元，后续治疗还需要xx元（家庭情况描述）。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.addImg')}}</p>
              <Upload ref="upload1" :on-success="ylImage" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.promise')}}</p>
              <i-input type="textarea"  v-model="ylDetails.promise" :rows="4" placeholder="1.所发起的个人大病求助项目，全部求助款项将用于求助人的治疗及康复费用，绝不挪用，如有剩余及时退还。
2.所发布的求助信息完整、准确、真实、及时、合法，无虚构或隐瞒情形。
3.将尊重赠与人的意愿，严格按照平台要求，及时说明求助人的治疗花费、家庭经济状况等平台要求说明的情况，并公示用。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.patientInformation')}}</p>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.patientName')}}</p>
                <i-input v-model="ylDetails.patientName" placeholder="请输入患者姓名"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.patientType')}}</p>
                <RadioGroup v-model="ylDetails.patientDocumentType">
                  <Radio label="1" value="1">{{$t('crowdfunding.idCard')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.passport')}}</Radio>
                </RadioGroup>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.idNumber')}}</p>
                <i-input v-model="ylDetails.patientDocumentNumber" placeholder="请输入证件号码"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.medicalInformation')}}</p>
                <Upload ref="upload2" :multiple="true" :on-success="ylImage2" :headers="uploadHeaders" :action="uploadUrl">
                  <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
                </Upload>
              </div>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" />{{$t('crowdfunding.sponsorInfo')}}</p>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorName')}}</p>
                <i-input v-model="ylDetails.sponsor" placeholder="请输入发起人的姓名"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorType')}}</p>
                <RadioGroup v-model="ylDetails.documentType">
                  <Radio label="1" value="1">{{$t('crowdfunding.idCard')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.passport')}}</Radio>
                </RadioGroup>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.idNumber')}}</p>
                <i-input v-model="ylDetails.documentNumber" placeholder="请输入证件号码"></i-input>
              </div>
            </div>
          </div>
           <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.estate')}}</p>
              <RadioGroup v-model="ylDetails.houseProperty">
                  <Radio label="无房产" value="无房产"></Radio>
                  <Radio label="一套房" value="一套房"></Radio>
                  <Radio label="两套房" value="两套房"></Radio>
                  <Radio label="三套以上" value="三套以上"></Radio>
              </RadioGroup>
            </div>
          </div>
          <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.carProduction')}}</p>
              <RadioGroup v-model="ylDetails.carProperty">
                  <Radio label="无车产" value="无车产"></Radio>
                  <Radio label="一辆车" value="一辆车"></Radio>
                  <Radio label="两辆车" value="两辆车"></Radio>
                  <Radio label="三辆以上" value="三辆以上"></Radio>
              </RadioGroup>
            </div>
          </div>
          <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.income')}}</p>
              <i-input v-model="ylDetails.income" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}/年</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.monetaryAssets')}}</p>
              <i-input v-model="ylDetails.property" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.arrears')}}</p>
              <i-input v-model="ylDetails.debt" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div style="width:300px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.otherApplications')}}</p>
              <RadioGroup v-model="ylDetails.otherFundType">
                  <Radio label="1" value="1">{{$t('crowdfunding.yes')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.no')}}</Radio>
              </RadioGroup>
            </div>
          </div>
        </div>
        <div class="window_btn" @click="ylClick()">{{$t('crowdfunding.submit')}}</div>
      </div>
    </div>
    <!--在线创业-->
    <div class="Initiate_mask2" v-if="Initiate_cy">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.basicInformation')}}</p>
          <img @click="showClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_tk_nr">
            <div style="width:350px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.demandAmount')}}</p>
              <i-input v-model="xyDetails.targetAmount" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
                  <span slot="append" style="color:red;margin-left:10px;">{{$t('crowdfunding.maxMoney')}} {{level}} {{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.fundraisingTitle')}}</p>
              <i-input v-model="xyDetails.fundingTitle" placeholder="一句话简单描述你的情况(25个字以内）"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.instructions')}}</p>
              <i-input v-model="xyDetails.instructions" type="textarea" :rows="4" placeholder="可为自己/家人/其他人发起该心愿众筹活动，
              请在此栏填写详细的自我介绍/再将心愿梦想完整准确的表述出来以获得金融家们的支持。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.addImg')}}</p>
              <Upload ref="upload1" :on-success="zcImage" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.sponsorInfo')}}</p>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorName')}}</p>
                <i-input v-model="xyDetails.sponsor" placeholder="请输入发起人的姓名"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorType')}}</p>
                <RadioGroup v-model="xyDetails.documentType">
                  <Radio label="1" value="1">{{$t('crowdfunding.idCard')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.passport')}}</Radio>
                </RadioGroup>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.idNumber')}}</p>
                <i-input v-model="xyDetails.documentNumber" placeholder="请输入证件号码"></i-input>
              </div>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.otherMaterials')}}</p>
              <Upload ref="upload2" :multiple="true" :on-success="zcImage2" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
        </div>
        <div class="window_btn" @click="xyClick(type = 2)">{{$t('crowdfunding.submit')}}</div>
      </div>
    </div>
    <!--线下申请-->
    <div class="Initiate_mask2" v-if="Initiate_xx">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" v-if="screenWidth > 1200" size="28" />
            <Icon type="ios-paper-outline" v-if="screenWidth < 1200" size="22" />{{$t('crowdfunding.applyPublicWelfare')}}</p>
          <img @click="xxClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_tk_nr">
            <div style="width:350px">
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.subsidyAmount')}}</p>
              <i-input v-model="xxDetails.applyAmount" placeholder="请输入金额" class="yanse">
                  <span slot="append">{{$t('crowdfunding.yuan')}}</span>
                  <span slot="append" style="color:red;margin-left:10px;">{{$t('crowdfunding.maxMoney')}} {{xx_level}} {{$t('crowdfunding.yuan')}}</span>
              </i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.publicWelfareTitle')}}</p>
              <i-input v-model="xxDetails.fundingTitle" placeholder="我为××地名××群体/个人送温暖行动"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.instructions')}}</p>
              <i-input v-model="xxDetails.instructions" type="textarea" placeholder="详细说明你与谁一起去到哪里为谁提供了帮助做了一件什么样公益活动，
              将整个活动从计划→实施→结束→感悟全部表述出来，并将活动的开销小票/照片等全流程照片上传。"></i-input>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.addImg')}}</p>
              <Upload ref="upload1" :on-success="xxImage" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.sponsorInfo')}}</p>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorName')}}</p>
                <i-input v-model="xxDetails.sponsor" placeholder="请输入发起人的姓名"></i-input>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.sponsorType')}}</p>
                <RadioGroup v-model="xxDetails.documentType">
                  <Radio label="1" value="1">{{$t('crowdfunding.idCard')}}</Radio>
                  <Radio label="2" value="2">{{$t('crowdfunding.passport')}}</Radio>
                </RadioGroup>
              </div>
              <div class="details" style="width:200px">
                <p><span>● </span>{{$t('crowdfunding.idNumber')}}</p>
                <i-input v-model="xxDetails.documentNumber" placeholder="请输入证件号码"></i-input>
              </div>
            </div>
          </div>
          <div class="window_tk_nr">
            <div>
              <p class="window_tk_nr_p"><Icon type="ios-navigate" size="20" color="#2987FE" /> {{$t('crowdfunding.otherMaterials')}}</p>
              <Upload ref="upload2" :multiple="true" :on-success="xxImage2" :headers="uploadHeaders" :action="uploadUrl">
                <Button style="color: #fff;" icon="ios-cloud-upload-outline">{{$t('uc.safe.upload')}}</Button>
              </Upload>
            </div>
          </div>
        </div>
        <div class="window_btn" @click="xxApplicationClick()">{{$t('crowdfunding.submit')}}</div>
      </div>
    </div>
  
    <!--在线详情-->
    <div class="Initiate_mask2" v-if="xinashang_details_show">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" size="28" v-if="screenWidth > 1200" />
          <Icon type="ios-paper-outline" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.details')}}</p>
          <img v-if="this.all == false" @click="xxClick()" src="../../assets/images/down.png" />
          <img v-if="this.all == true" @click="all_xxClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_xx_name2">
             <div class="window_tx_left">
                <img class="window_tx" src="../../assets/images/activity_mobile.jpg" />
                <p>{{xs_details.sponsor}}</p>
                <p>{{$t('crowdfunding.initiate')}}</p>
             </div>
             <div class="window_tx_right" v-if="xs_details.type == 1"><Icon type="ios-heart" size="27" /> {{$t('crowdfunding.wish')}}</div>
             <div class="window_tx_right" style="background: #0071FE;" v-if="xs_details.type == 3"><Icon type="ios-calendar" size="27" /> {{$t('crowdfunding.medical')}}</div>
             <div class="window_tx_right" style="background: #0AC159;" v-if="xs_details.type == 2"><Icon type="md-medkit" size="27" /> {{$t('crowdfunding.pioneer')}}</div>
             <div style="clear:both;"></div>
          </div>
          <div class="window_xx_title">
            <p>{{xs_details.fundingTitle}}</p>
            <p>{{dateFormat(xs_details.passTime)}}—{{dateFormat(xs_details.endTime)}}</p>
          </div>
          <div class="window_xs_content">
            <Row>
                <Col span="8" class="window_col">
                  <p>{{xs_details.targetAmount}}</p>
                  <p>{{$t('crowdfunding.targetAmount')}}(USDT)</p>
                </Col>
                <Col span="8" class="window_col">
                  <p>{{xs_details.amountReceived}}</p>
                  <p>{{$t('crowdfunding.amountRaised')}}(USDT)</p>
                </Col>
                <Col span="8" class="window_col">
                  <p>{{xs_details.times}}</p>
                  <p>{{$t('crowdfunding.helpTimes')}}({{$t('crowdfunding.times')}})</p>
                </Col>
            </Row>
            <div class="window_xs_lable">
              <div class="window_xs_lable_title">
                <p class="left"><Icon type="ios-paper" size="28" v-if="screenWidth > 1200" />
                <Icon type="ios-paper" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.donationRecord')}}</p>
                <!-- <p class="right">查看全部></p> -->
                <p style="clear:both;"></p>
              </div>
              <div class="window_xs_lable_content" style=" height:130px;overflow-y: auto;">
                <div v-for="(item, index) in recordList" :key="index">
                  <p>{{item.memberName.replace(/^(\d{1})\d{9}(\d{1})$/,'$1*********$2')}}</p>
                  <p>{{item.fundingMoney}} USDT</p>
                  <p>{{dateFormat(item.fundingTime)}}</p>
                </div>
              </div>
            </div>
            <div style="width:100%;height:20px;"></div>
          </div>
          <div style="width:95%;margin:0 auto; margin-top:20px;">
            <Tabs class="tabs2">
              <TabPane :label="$t('crowdfunding.projectStory')" icon="ios-bookmarks" style=" height:330px;overflow-y: auto;">
                <div class="window_xx_content" style="margin-top:20px;">
                  <p>
                    {{xs_details.instructions}}
                  </p>
                  <div class="window_xx_img"  v-for="(item, index) in xs_details_img" :key="index">
                    <img style="float:left;margin-right:15px;" :src="item.picUrl"/>
                  </div>
                  <div style="clear:both;"></div>
                </div>
              </TabPane>
              <TabPane :label="$t('crowdfunding.comment')" icon="md-text" style=" height:330px;overflow-y: auto;">
                <div style="width:95%;margin:0 auto;"  v-for="(item, index) in commentList" :key="index">
                  <div class="window_xs_pl">
                    <p>
                      {{item.memberName.replace(/^(\d{1})\d{9}(\d{1})$/,'$1*********$2')}}
                    </p>
                    <p>{{dateFormat(item.commonTime)}}</p>
                  </div>
                  <div class="window_xs_pl_con">
                    <p>{{item.common}}</p>
                  </div>
                </div>
              </TabPane>
            </Tabs>
          </div>
        </div>
        <div style="display:flex;">
          <div class="window_btn2" @click="PLClick()">{{$t('crowdfunding.comment2')}}</div>
          <div class="window_btn2" @click="JKClick()">{{$t('crowdfunding.donate2')}}</div>
        </div>
      </div>
    </div>
    <!--线下详情-->
    <div class="Initiate_mask2" v-if="xx_details_show">
      <div class="window_tk">
        <div class="window_tk_title">
          <p><Icon type="ios-paper-outline" v-if="screenWidth > 1200" size="28" />
          <Icon type="ios-paper-outline" v-if="screenWidth < 1200" size="22" />{{$t('crowdfunding.details')}}</p>
          <img v-if="this.all == false" @click="xxClick()" src="../../assets/images/down.png" />
          <img v-if="this.all == true" @click="all_xxClick2()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content">
          <div class="window_xx_name">
             <img class="window_tx" src="../../assets/images/activity_mobile.jpg" />
             <p>{{xx_details.sponsor}}</p>
             <p>{{$t('crowdfunding.initiate')}}</p>
          </div>
          <div class="window_xx_title">
            <p>{{xx_details.fundingTitle}}</p>
            <p>{{dateFormat(xx_details.passTime)}}</p>
          </div>
          <div class="window_xx_content">
            <p>{{xx_details.instructions}}</p>
            <div class="window_xx_img" v-for="(item, index) in getWelfareDetail" :key="index">
              <img style="float:left;margin-right:15px;" :src="item.picUrl"/>
            </div>
            <div style="clear:both;"></div>
          </div>
        </div>
        <div class="window_btn2" v-if="isLogin" @click="xx_apply()">{{$t('crowdfunding.applySubsidy')}}</div>
      </div>
    </div>
    <!--我要捐款-->
    <div class="Initiate_mask3" v-if="JK_details_show">
      <div class="window_tk2">
        <div class="window_tk_title">
          <p><Icon type="logo-usd" size="28" v-if="screenWidth > 1200" />
          <Icon type="logo-usd" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.donationMoney')}}</p>
          <img @click="JKxxClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content2">
          <div style="width:90%;margin:0 auto">
          <Row :gutter="10">
              <i-col :xs="{ span: 8 }" :lg="{ span: 8}">
                  <div class="window_col2" :class="jkMoney == 1 ? 'window_col3': ''" @click="jkMoney1()">
                    <p>50 USDT</p>
                  </div>
              </i-col>
              <i-col :xs="{ span: 8 }" :lg="{ span: 8}">
                  <div class="window_col2" :class="jkMoney == 2 ? 'window_col3': ''" @click="jkMoney2()">
                    <p>100 USDT</p>
                  </div>
              </i-col>
              <i-col :xs="{ span: 8 }" :lg="{ span: 8}">
                  <div class="window_col2" :class="jkMoney == 3 ? 'window_col3': ''" @click="jkMoney3()">
                    <p>300 USDT</p>
                  </div>
              </i-col>
          </Row>
          </div>
          <div class="window_tk_nr">
              <div class="window_jkjg" @click="jkMoney4()">
                <i-input v-model="fundingMoney" :placeholder="$t('crowdfunding.otherAmount')" class="yanse2">
                    <span slot="append">USDT</span>
                </i-input>
              </div>
          </div>
        </div>
        <div style="color:red;width:100%;text-align:center;position: relative;top: -60px;">{{$t('crowdfunding.rate')}}</div>
        <div v-if="this.xs_details.targetAmount * 0.1 >= this.fundingMoney">
          <div class="window_btn2" v-if="this.zje >= this.fundingMoney" @click="JKxxClick2()">{{$t('crowdfunding.donate')}}</div>
          <div class="window_jk_title" v-else>{{$t('crowdfunding.lessThan')}}{{zje}}</div>
        </div>
        <div class="window_jk_title" v-else>{{$t('crowdfunding.lessThan')}}{{xs_details.targetAmount * 0.1}}</div>
      </div>
    </div>
    <!--评论-->
    <div class="Initiate_mask3" v-if="PL_details_show">
      <div class="window_tk2">
        <div class="window_tk_title">
          <p><Icon type="logo-create" size="28" v-if="screenWidth > 1200" />
          <Icon type="logo-create" size="22" v-if="screenWidth < 1200" />{{$t('crowdfunding.comment')}}</p>
          <img @click="PL_showClick()" src="../../assets/images/down.png" />
        </div>
        <div class="window_tk_content2">
          <div class="window_tk_nr">
              <div class="window_jkjg2">
                <i-input v-model="common" type="textarea" :rows="10" :placeholder="$t('crowdfunding.comment')"></i-input>
              </div>
          </div>
        </div>
        <div class="window_btn2" @click="PLxxClick()">{{$t('crowdfunding.comment')}}</div>
      </div>
    </div>
  </div>
</template>

<script>
import Swiper from 'swiper';
var moment = require("moment");
export default {
  name: "Slider",
  data() {
    return {
      uploadHeaders:{'x-auth-token':localStorage.getItem('TOKEN')},
      uploadUrl:this.host+'/uc/upload/oss/image',
      money:'',
      totalMoney:'',
      totalProject:'',
      totalTimes:'',
      isLogin:false,
      mask:false,
      all:false,
      all_show:false,
      all_show2:false,
      Initiate_select:false,
      Initiate_xy:false,
      Initiate_yl:false,
      Initiate_yl2:false,
      Initiate_cy:false,
      Initiate_xx:false,
      xx_details_show:false,
      xinashang_details_show:false,
      JK_details_show:false,
      PL_details_show:false,
      screenWidth: document.body.clientWidth,
      level:'',
      item:'',
      xx_level:'',
      code:'',
      ifxsBtn:'',
      ifxxBtn:'',
      zje:'',
      list:[],
      xs_details:{},
      xs_details_img:{},
      xx_details:{},
      recordList:[],
      commentList:[],
      enablewelfare:[],
      getWelfareDetail:[],
      ylDetails:{
        targetAmount:"",
        fundingTitle:"",
        fundingUse:"",
        plan:"",
        story:"",
        picUrl:"",
        promise:"",
        patientName:"",
        patientDocumentType:"1",
        patientDocumentNumber:"",
        visitInformation:'',
        crowdPics:{
          pics:[]
        },
        sponsor:"",
        documentType:"1",
        documentNumber:"",
        houseProperty:"无房产",
        carProperty:"无车产",
        income:"",
        property:"",
        debt:"",
        otherFundType:"1",
      },
      xyDetails:{
        targetAmount:"",
        fundingTitle:"",
        instructions:"",
        picUrl:"",
        sponsor:"",
        documentType:"1",
        documentNumber:"",
        crowdPics:{
          pics:[]
        },
      },
      xxDetails:{
        applyAmount:"",
        fundingTitle:"",
        instructions:"",
        picUrl:"",
        sponsor:"",
        documentType:"1",
        documentNumber:"",
        material:"",
        crowdPics:{
          pics:[]
        },
      },
      fundingMoney:"",
      jkMoney:"",
      common:""
    };
  },
  computed: {
    isLogin: function() {
      return this.$store.getters.isLogin;
      console.log(this.$store.getters.isLogin)
    }
  },
  created: function() {
    this.init();
    this.fmoney();
    this.isLogin = this.$store.getters.isLogin;
    console.log(this.isLogin)
    this.info();
  },
  filters:{
    dateFormat: function(tick) {
      return moment(tick).format("YYYY-MM-DD");
    },
  },
  computed: {
    lang() {
      return this.$store.state.lang;
    },
    langPram(){
      if(this.$store.state.lang == "简体中文"){
        return "CN";
      }
      if(this.$store.state.lang == "English"){
        return "EN";
      }
      return "CN";
    }
  },
  methods: {
    dateFormat: function(tick) {
      return moment(tick).format("YYYY-MM-DD");
    },
    ylImage (res, file,fileList) {
      this.$refs.upload1.fileList=[fileList[fileList.length-1]];
      this.ylDetails.picUrl=res.data;
      console.log(this.ylDetails.picUrl)
    },
    ylImage2 (res, file,fileList) {
      this.ylDetails.visitInformation = res.data
      let picUrl = {}
      picUrl.picUrl = res.data
      this.ylDetails.crowdPics.pics.push(picUrl)
      console.log(this.ylDetails.crowdPics)
    },
    zcImage (res, file,fileList) {
      this.$refs.upload1.fileList=[fileList[fileList.length-1]];
      this.xyDetails.picUrl=res.data;
             console.log(res)
    },
    zcImage2 (res, file,fileList) {
      let picUrl = {}
      picUrl.picUrl = res.data
      this.xyDetails.crowdPics.pics.push(picUrl)
      console.log(this.xyDetails.crowdPics)
    },
    xxImage (res, file,fileList) {
        this.$refs.upload1.fileList=[fileList[fileList.length-1]];
          this.xxDetails.picUrl=res.data;
             console.log(this.xxDetails.picUrl)
    },
    xxImage2 (res, file,fileList) {
      let picUrl = {}
      picUrl.picUrl = res.data
      this.xxDetails.crowdPics.pics.push(picUrl)
      console.log(this.xyDetails.crowdPics)
    },
    init() {
      this.$store.commit("navigate", "nav-crowdfunding");
    },
    syadd_show(){
      this.all_show = true;
    },
    allClick(){
      this.all_show = false;
    },
    syadd_show2(){
      this.all_show2 = true;
    },
    allClick2(){
      this.all_show2 = false;
    },
    maskClick() {
      this.mask = true;
      this.Initiate_select = true;
    },
    maskClick2() {
      this.mask = false;
      this.Initiate_select = false;
    },
    selectClick1(){
      this.Initiate_select = false;
      this.Initiate_xy = true;
    },
    selectClick2(){
      this.Initiate_select = false;
      this.Initiate_yl = true;
    },
    selectClick3(){
      this.Initiate_select = false;
      this.Initiate_cy = true;
    },
    showClick(){
      this.mask = false;
      this.Initiate_xy = false;
      this.Initiate_yl = false;
      this.Initiate_cy = false;
    },
    xyClick(type){
      var that = this;
      if(this.xyDetails.targetAmount == ""){
        this.$Message.error('请输入目标金额');
        return false
      };
      if(this.xyDetails.targetAmount > this.level){
        this.$Message.error('目标金额不能大于' + this.level);
        return false
      };
      if(this.xyDetails.fundingTitle == ""){
        this.$Message.error('请输入标题');
        return false
      };
      if(this.xyDetails.instructions == ""){
        this.$Message.error('请输入说明');
        return false
      };
      if(this.xyDetails.picUrl == ""){
        this.$Message.error('请添加图片');
        return false
      };
      if(this.xyDetails.sponsor == ""){
        this.$Message.error('请输入发起人');
        return false
      };
      if(this.xyDetails.documentNumber == ""){
        this.$Message.error('请输入证件号');
        return false
      };
      let param = {}
          param.funding = {}
          param.funding.targetAmount = this.xyDetails.targetAmount //目标金额
          param.funding.type = type
          param.funding.fundingTitle = this.xyDetails.fundingTitle //标题
          param.funding.instructions = this.xyDetails.instructions //说明
          param.funding.picUrl = this.xyDetails.picUrl //添加图片
          param.funding.sponsor = this.xyDetails.sponsor //发起人
          param.funding.documentType = this.xyDetails.documentType //发起人证件类型
          param.funding.documentNumber = this.xyDetails.documentNumber //证件号
          param.crowdPics = this.xyDetails.crowdPics //材料地址
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.wish,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.mask = false;
            this.Initiate_xy = false;
            this.Initiate_cy = false;
            this.$Message.success(resp.message);
          } else {
            this.$Message.error(resp.message);
          }
        });
      
    },
    ylClick(){
      var that = this;
      if(this.ylDetails.targetAmount == ""){
        this.$Message.error('请输入目标金额');
        return false
      };
      if(this.ylDetails.targetAmount > this.level){
        this.$Message.error('目标金额不能大于' + this.level);
        return false
      };
      if(this.ylDetails.fundingTitle == ""){
        this.$Message.error('请输入标题');
        return false
      };
      if(this.ylDetails.fundingUse == ""){
        this.$Message.error('请输入用途');
        return false
      };
      if(this.ylDetails.plan == ""){
        this.$Message.error('请输入计划');
        return false
      };
      if(this.ylDetails.story == ""){
        this.$Message.error('请输入故事');
        return false
      };
      if(this.ylDetails.picUrl == ""){
        this.$Message.error('请添加图片');
        return false
      };
      if(this.ylDetails.promise == ""){
        this.$Message.error('请输入承诺');
        return false
      };
      if(this.ylDetails.patientName == ""){
        this.$Message.error('请输入患者姓名');
        return false
      };
      if(this.ylDetails.patientDocumentNumber == ""){
        this.$Message.error('请输入患者证件号');
        return false
      };
      if(this.ylDetails.visitInformation == ""){
        this.$Message.error('请选择就诊信息图片');
        return false
      };
      if(this.ylDetails.sponsor == ""){
        this.$Message.error('请输入发起人');
        return false
      };
      if(this.ylDetails.documentNumber == ""){
        this.$Message.error('请输入发起人证件号');
        return false
      };
      if(this.ylDetails.income == ""){
        this.$Message.error('请输入患者收入情况');
        return false
      };
      if(this.ylDetails.property == ""){
        this.$Message.error('请输入资产情况');
        return false
      };
      if(this.ylDetails.debt == ""){
        this.$Message.error('请输入欠款情况');
        return false
      };
      let param = {}
          param.funding = {}
          param.funding.targetAmount = this.ylDetails.targetAmount //目标金额
          param.funding.type = 3
          param.funding.fundingTitle = this.ylDetails.fundingTitle //标题
          param.funding.fundingUse = this.ylDetails.fundingUse //用途
          param.funding.plan = this.ylDetails.plan //计划
          param.funding.story = this.ylDetails.story //故事
          param.funding.picUrl = this.ylDetails.picUrl //添加图片
          param.funding.promise = this.ylDetails.promise //承诺
          param.funding.patientName = this.ylDetails.patientName //患者姓名
          param.funding.patientDocumentType = this.ylDetails.patientDocumentType //证件类型
          param.funding.patientDocumentNumber = this.ylDetails.patientDocumentNumber //证件号
          param.funding.sponsor = this.ylDetails.sponsor //发起人
          param.funding.documentType = this.ylDetails.documentType //发起人证件类型
          param.funding.documentNumber = this.ylDetails.documentNumber //证件号
          param.funding.houseProperty = this.ylDetails.houseProperty //房屋
          param.funding.carProperty = this.ylDetails.carProperty //车辆
          param.funding.income = this.ylDetails.income //患者收入情况
          param.funding.property = this.ylDetails.property //资产情况
          param.funding.debt = this.ylDetails.debt //欠款情况
          param.funding.otherFundType = this.ylDetails.otherFundType //是否其他平台发起筹款
          param.crowdPics = this.ylDetails.crowdPics 
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.medical,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.mask = false;
            this.Initiate_yl = false;
            this.$Message.success(resp.message);
          } else {
            this.$Message.error(resp.message);
          }
        });
    },
    xx_apply(){
      this.mask = true;
      this.Initiate_xx = true;
      this.xx_details_show = false;
    },
    xxApplicationClick(){
      var that = this;
      if(this.xxDetails.applyAmount == ""){
        this.$Message.error('请输入目标金额');
        return false
      };
      if(this.xxDetails.applyAmount > this.xx_level){
        this.$Message.error('目标金额不能大于' + this.xx_level);
        return false
      };
      if(this.xxDetails.targetAmount > this.xxDetails){

      };
      if(this.xxDetails.fundingTitle == ""){
        this.$Message.error('请输入标题');
        return false
      };
      if(this.xxDetails.instructions == ""){
        this.$Message.error('请输入说明');
        return false
      };
      if(this.xxDetails.picUrl == ""){
        this.$Message.error('请添加图片');
        return false
      };
      if(this.xxDetails.sponsor == ""){
        this.$Message.error('请输入发起人');
        return false
      };
      if(this.xxDetails.documentNumber == ""){
        this.$Message.error('请输入证件号');
        return false
      };
      let param = {}
          param.crowdWelfare = {}
          param.crowdWelfare.applyAmount = this.xxDetails.applyAmount //目标金额
          param.crowdWelfare.fundingTitle = this.xxDetails.fundingTitle //标题
          param.crowdWelfare.instructions = this.xxDetails.instructions //说明
          param.crowdWelfare.picUrl = this.xxDetails.picUrl //添加图片
          param.crowdWelfare.sponsor = this.xxDetails.sponsor //发起人
          param.crowdWelfare.documentType = this.xxDetails.documentType //发起人证件类型
          param.crowdWelfare.documentNumber = this.xxDetails.documentNumber //证件号
          param.crowdPics = this.xxDetails.crowdPics //材料上传
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.publicWelfare,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.mask = false;
            this.Initiate_xx = false;
            this.$Message.success(resp.message);
          } else {
            this.$Message.error(resp.message);
          }
        });
    },
    xxClick(){
      this.mask = false;
      this.Initiate_xx = false;
      this.xx_details_show = false;
      this.xinashang_details_show = false;
    },
    all_xxClick(){
      this.mask = false;
      this.xinashang_details_show = false;
      this.all_show = true;
      this.all = false;
    },
    all_xxClick2(){
      this.mask = false;
      this.xx_details_show = false;
      this.all_show2 = true;
      this.all = false;
    },
    xx_detailsClick(item){
      this.mask = true;
      this.xx_details_show = true;
      this.xx_details = item;
       let param = {}
          param.id = item.id
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.getWelfareDetail,param)
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            console.log(response);
            this.getWelfareDetail = resp.data.picsUrl
          }
        });
    },
    xx_detailsClick2(item){
      this.mask = true;
      this.xx_details_show = true;
      this.all_show2 = false;
      this.all = true;
      this.xx_details = item;
    },
    xs_details_show(item){
      this.item = item;
      this.mask = true;
      this.xinashang_details_show = true;
      this.xs_details = item;
      this.donateList(item)
      console.log(this.xs_details)
    },
    xs_details_show2(item){
      this.item = item;
      this.mask = true;
      this.xinashang_details_show = true;
      this.all_show = false;
      this.all = true;
      this.xs_details = item;
      this.donateList(item)
    },
    JKClick(){
      this.JK_details_show = true;
      this.zje = this.xs_details.targetAmount - this.xs_details.amountReceived;
    },
    JKxxClick(){
      this.JK_details_show = false;
    },
    jkMoney1(){
      this.jkMoney = 1;
      this.fundingMoney = 50;
    },
    jkMoney2(){
      this.jkMoney = 2;
      this.fundingMoney = 100;
    },
    jkMoney3(){
      this.jkMoney = 3;
      this.fundingMoney = 300;
    },
    jkMoney4(){
      this.jkMoney = 0;
      this.fundingMoney = '';
    },
    JKxxClick2(){
      var that = this;
      if(this.fundingMoney == ""){
        this.$Message.error('请选择/输入捐款金额');
        return false
      };
      if(this.fundingMoney < 1){
        this.$Message.error('最少捐1USDT');
        return false
      };
      let param = {}
          param.memberName = "" 
          param.fundingType = this.xs_details.type
          param.fundingId = this.xs_details.id
          param.fundingMoney = this.fundingMoney
          param.fundingTitle = this.xs_details.fundingTitle
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.donate,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.JK_details_show = false;
            this.$Message.success(resp.message);
            this.totalMoney = resp.data.totalMoney;
            this.totalProject = resp.data.totalProject;
            this.totalTimes = resp.data.totalTimes;
            this.fmoney();
            this.fmoney2();
            this.fmoney3();
            this.donateList(this.item)
          } else {
            this.$Message.error('捐款失败!' + resp.message);
          }
        });
    },
    PLClick(){
      this.PL_details_show = true;
    },
    PL_showClick(){
      this.PL_details_show = false;
    },
    PLxxClick(){
      var that = this;
      if(this.common == ""){
        this.$Message.error('请输入评论后再提交');
        return false
      };
      let param = {}
          param.memberName = "" 
          param.fundingType = this.xs_details.type
          param.fundingId = this.xs_details.id
          param.common = this.common
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.comment,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.PL_details_show = false;
            this.$Message.success('评论成功!请等待后台审核');
            this.donateList(this.item)
          } else {
            this.$Message.error('评论失败!');
          }
        });
    },
    initSwiper() {
      if(this.screenWidth > 1200){
      var mySwiper = new Swiper('.swiper-container',{
          autoplay: {
            delay: 3000,
            disableOnInteraction: false,
          },
          loop: true,
          speed:1000,
          autoplayDisableOnInteraction : false,
          centeredSlides: true,    
          paginationClickable: true,
          slidesPerView: 3,
          navigation: {
            nextEl: ".swiper-button-next",
            prevEl: ".swiper-button-prev",
          },
          onInit:function(swiper){
		      	swiper.slides[3].className="swiper-slide swiper-slide-active";//第一次打开不要动画
			    },
          breakpoints: { 
            668: {
              slidesPerView: 3,
            }
          }
      }) 
      } 
      if(this.screenWidth < 1200){
        var mySwiper = new Swiper('.swiper-container',{
          autoplay: {
            delay: 3000,
            disableOnInteraction: false,
          },
          loop: true,
          speed:1000,
          autoplayDisableOnInteraction : false,
          centeredSlides: true,    
          paginationClickable: true,
          slidesPerView: 1,
          navigation: {
            nextEl: ".swiper-button-next",
            prevEl: ".swiper-button-prev",
          },
          onInit:function(swiper){
		      	swiper.slides[1].className="swiper-slide swiper-slide-active";//第一次打开不要动画
			    },
          breakpoints: { 
            668: {
              slidesPerView: 1,
            }
          }
      }) 
      }
    },
    doResize(){
        setTimeout(function(){
            //手动触发窗口resize事件
            if(document.createEvent) {
                var event = document.createEvent("HTMLEvents");
                event.initEvent("resize", true, true);
                window.dispatchEvent(event);
            } else if(document.createEventObject) {
                window.fireEvent("onresize");
            }
        },100);
    },
    info() {
      var that = this;
      if (this.isLogin == true) {
          this.getLevelRight();
      }
      this.$http
        .get(this.host + this.api.crowdfunding.totalAmount)
        .then(response => {
          var resp = response.body.data;
          this.totalMoney = resp.totalMoney;
          this.totalProject = resp.totalProject;
          this.totalTimes = resp.totalTimes;
          this.fmoney();
          this.fmoney2();
          this.fmoney3();
        });
      this.$http
        .get(this.host + this.api.crowdfunding.crowdfundingList)
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            response.body.data.funding.forEach((v, k) => {
              this.list.push(v)
            })
            response.body.data.medicalfunding.forEach((v, k) => {
              this.list.push(v)
            })
            this.$nextTick(()=>{
              this.initSwiper()
              this.doResize();
            })
          } 
        });
        this.$http
        .get(this.host + this.api.crowdfunding.offlineList)
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            console.log(response);
            this.enablewelfare = resp.data
          }
        });
    },
    fmoney() {
      var	n = 2;
      var s = parseFloat((this.totalMoney + "").replace(/[^\d\.-]/g, "")).toFixed(n) + "";
      console.log(s)
    	var l = s.split(".")[0].split("").reverse(), r = s.split(".")[1];
      var t = "";
      console.log(l)
    	for (var i = 0; i < l.length; i++) {
	    	t += l[i] + ((i + 1) % 3 == 0 && (i + 1) != l.length ? "," : "");
	    }
	    return this.totalMoney = t.split("").reverse().join("") + "." + r;
    },
    fmoney2() {
      var	n = 2;
      var s = parseFloat((this.totalProject + "").replace(/[^\d\.-]/g, "")).toFixed(n) + "";
      console.log(s)
    	var l = s.split(".")[0].split("").reverse(), r = s.split(".")[1];
      var t = "";
      console.log(l)
    	for (var i = 0; i < l.length; i++) {
	    	t += l[i] + ((i + 1) % 3 == 0 && (i + 1) != l.length ? "," : "");
	    }
	    return this.totalProject = t.split("").reverse().join("");
    },
    fmoney3() {
      var	n = 2;
      var s = parseFloat((this.totalTimes + "").replace(/[^\d\.-]/g, "")).toFixed(n) + "";
      console.log(s)
    	var l = s.split(".")[0].split("").reverse(), r = s.split(".")[1];
      var t = "";
      console.log(l)
    	for (var i = 0; i < l.length; i++) {
	    	t += l[i] + ((i + 1) % 3 == 0 && (i + 1) != l.length ? "," : "");
	    }
	    return this.totalTimes = t.split("").reverse().join("");
    },
    getLevelRight() {
      this.$http
        .post(this.host + this.api.crowdfunding.level)
        .then(response => {
          var resp = response.body;
          console.log(response)
          this.ifxsBtn = resp.data.crowdRight;
          this.ifxxBtn = resp.data.welfareRight;
          console.log(this.code)
          this.level = resp.data.crowdRight * 10000;
          this.xx_level = resp.data.welfareRight * 10000;
        });
    },
    donateList(item){
      var that = this;
      let param = {}
          param.id = item.id
          param.type = item.type
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.donateList,param,{emulateJSON:false})
        .then(response => {
          var resp = response.body;
            console.log(resp);
            this.recordList = resp.data.recordList;
            this.commentList = resp.data.fundingCommonList;
            this.xs_details = resp.data.funding;
            this.xs_details_img = resp.data.picsUrl;
        });
    }
  }
};
</script>
<style lang="scss" scoped>

    .all_xs{
      width: 90%;
      margin: 0 auto;
      height: auto;
      padding-top: 30px;
      border-bottom: 1px solid #ccc;
    }
    .all_xs_img{
      float: left;
    }
    .all_xs_img img{
      width: 300px;
      height: 150px;
      border-radius: 10px;
    }
    .all_xs_text{
      float: right;
      width: calc(100% - 300px);
      color: #000;
      border-radius: 10px;
    }
    .all_xs_text_left{
      float: left;
      margin-left: 20px;
    }
    .all_xs_text_title{
      width: 300px;
      font-size: 22px;
      overflow:hidden;
      text-overflow:ellipsis;
      white-space:nowrap;
    }
    .all_xs_text_content{
      background: #f5f5f5;
      border-radius: 10px;
      padding: 0px 10px;
      display: flex;
    }
    .all_xs_text_content>div{
      width: auto;
      text-align: center;
    }
    .all_xs_text_content p{
      line-height: 30px;
      min-width: 100px;
      padding: 5px 10px;
      text-align: center;
    }
    .all_xs_text_content p:nth-child(2){
      color: red;
      font-size: 20px;
      font-weight: bold;
    }
    .all_xs_text_content p:nth-child(1){
      color: #000;
      font-size: 16px;
      white-space:nowrap;
    }
    .all_xs_text_money{
      position: relative;
      top: 25px;
      color: #FE6700;
    }
    .all_xs_text_right{
      float: right;
      position: relative;
    }
    .all_xs_right_top{
      font-size: 16px;
      height: 40px;
      line-height: 40px;
    }
    .all_xs_right_top span:nth-child(1){
      color: red;
      font-size: 20px;
    }
    .all_xs_right_top span:nth-child(2){
      color: #666;
      font-size: 14px;
      position: relative;
      top:0px;
    }
    .all_xs_right_bottom{
      margin-top: 80px;
      font-size: 14px;
    }

    ::v-deep.ivu-radio-group{
      .ivu-radio-wrapper{
        color: #000;
        .ivu-radio-checked {
          .ivu-radio-inner {
            border-color: #2d8cf0;
          }
          .ivu-radio-inner:after {
            background: #2d8cf0;
          }
        }
      }
    }
    ::v-deep.ivu-upload{
      .ivu-btn{
        .ivu-btn-default{
          background-color: #f5f5f5;
        }
      }
    }
    ::v-deep.ivu-input-wrapper{
      .ivu-input{
        background-color: #f5f5f5;
        color: #000;
        border-color: #f5f5f5;
      }
    }
    ::v-deep.yanse{
      .ivu-input{
        color: #FD4E17!important;
      }
    }
    ::v-deep.yanse2{
      position: relative;
      top: 30px;
      .ivu-input{
        color: #FD4E17!important;
        background-color: #fff;
        color: #000;
        height: 45px;
        border-color: #f5f5f5;
        box-shadow: 0px 0px 4px #ccc;
        border-radius: 10px;
      }
    }
    ::v-deep.ivu-icon{
      position: relative;
      top: -1px;
    }
    ::v-deep.ivu-input-group {
      .ivu-input-group-append{
        background-color: #fff;
        border-bottom: 1px solid #fff;
        color: #000;
        border: 1px solid #fff;
      }
    }
    ::v-deep.ivu-card{
      cursor: pointer;
      border-radius: 30px;
      border: 2px solid #fff;
    }
    ::v-deep.ivu-card:hover{
      border: 2px solid #0071FE;
      position: relative;
      top: -15px;
    }
    ::v-deep.tabs1{
      .ivu-tabs-bar{
        .ivu-tabs-nav-container{
          .ivu-tabs-nav-wrap{
            .ivu-tabs-nav-scroll{
              .ivu-tabs-nav{
                .ivu-tabs-tab{
                  font-size: 26px;
                }
                .ivu-icon{
                  width: 24px;
                  height: 24px;
                }
              }
            }
          }
        }
      }
    }
    ::v-deep.tabs2{
      .ivu-tabs-bar{
        .ivu-tabs-nav-container{
          .ivu-tabs-nav-wrap{
            .ivu-tabs-nav-scroll{
              .ivu-tabs-nav{
                .ivu-tabs-tab{
                  font-size: 26px;
                }
                .ivu-icon{
                  width: 24px;
                  height: 24px;
                }
                .ivu-tabs-ink-bar{
                  width: 175px!important;
                }
              }
            }
          }
        }
      }
    }
   /*swiper*/
    .swiper-wrapper{
      margin: 20px 0px;
    }
    .swiper-slide {
    	-webkit-transition: transform 1.0s;
    	-moz-transition: transform 1.0s;
    	-ms-transition: transform 1.0s;
    	-o-transition: transform 1.0s;
    	-webkit-transform: scale(0.7);
    	transform: scale(0.7);
      cursor: pointer;
    }
    .swiper-slide-active,.swiper-slide-duplicate-active {
    	-webkit-transform: scale(1);
    	transform: scale(1);
    }
    .swiper-slide a {
    	background: #fff;
    	padding:10px;
    	display: block;
    	border-radius: 14px;
    }
    .swiper-slide>div {
    	width: auto;
      box-shadow: 0px 0px 4px #ccc;
      border-radius: 14px;
      height: 300px;
    }
    .swiper-slide img {
    	width: 100%;
      border-top-left-radius: 14px;
      border-top-right-radius: 14px;
    	display: block;
      height: 194px;
    }
    .swiper-slide2 img {
    	width: 100%;
      border-top-left-radius: 14px;
      border-top-right-radius: 14px;
    	display: block;
      height: 254px;
    }
    .swiper-button-next,.swiper-button-prev{
      outline: none;
      opacity: 0.6;
      width: 25px;
      height: 40px;
    }
    .swiper_title{
      width: 90%;
      margin: 0 auto;
      display: flex;
      margin-top:5px;
    }
    .swiper_title p{
      width:50%;
    }
    .swiper_title p:nth-child(1){
      color:#373737;
      font-size: 20px;
      font-weight: 600;
      text-align: left;
overflow:hidden;
text-overflow:ellipsis;
white-space:nowrap;
    }
    .swiper_title p:nth-child(2){
      color:#373737;
      font-size: 12px;
      position: relative;
       text-align: right;
      top: 6px;
    }
    .swiper_content{
      margin-top: 10px;
    }
    .swiper_content>div p:nth-child(1){
      color:#FE6700;
      font-size: 20px;
    }
    .swiper_content>div p:nth-child(2){
      color:#707070;
      font-size: 13px;
    }
    /*swiper 结束*/
  .activity {
     width: 100%;
    background: rgba(242,246,250,1) !important;
    height: 100%;
    background-size: cover;
    position: relative;
    overflow: hidden;
    padding-bottom: 20px;
    padding-top: 60px;
    color: #fff;
  }
  .activity .banner{
      width:100%;
  }
  .banner_text{
    position: absolute;
    top: 60px;
    width: 100%;
    text-align: center;
    color: #fff;
    font-weight: 500;
    font-size: 28px;
    margin-top: 40px;
  }
  .money{
    font-size: 64px;
    color: red;
    letter-spacing:14px;
  }
  .money_span{
    font-size: 33px;
    color: #fff;
    font-weight: 400;
    position: relative;
    top: -10px;
  }
  .banner_content{
    width: 1200px;
    margin: 0 auto;
  }
  .banner_content p{
    font-size: 20px;
    letter-spacing:4px;
  }
  .banner_content p span{
    font-size: 18px;
    letter-spacing:1px;
    position: relative;
    top:-1px;
  }
  .banner_content p:nth-child(1){
    float: left;
    font-weight: 100;
  }
  .banner_content p:nth-child(2){
    float: right;
    font-weight: 100;
  }
  .br{
    display: none;
  }
  .tab{
    width: 1200px;
    margin: 0 auto;
    font-weight: 500;
  }
  .tab_all{
    position: relative;
    top: -42px;
    float: right;
    color: #A1A1A1;
    cursor: pointer;
  }
  .tab_all2{
    width: 100%;
    text-align: center;
    font-size: 18px;
    margin-top: 10px;
    display: none;
  }
  .btn{
    margin: 0 auto;
    margin-top: 10px;
    width:240px;
    height: 60px;
    line-height: 60px;
    font-size: 23px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .btn span{
    cursor: pointer;
  }
  .bottom{
    width: 1200px;
    text-align: center;
    margin-top: 42px;
  }
  .bottom_text{
    width: 1200px;
    text-align: left;
    color: #2987FE;
  }
  .mask{
    position:absolute;left:0px;top:0px;
    background:rgba(0, 0, 0, 0.5);
    width:100%;
    height:100%;
    z-Index:1; 
  }
  .Initiate_mask{
    position: fixed;
    top: 20%;
    width: 100%;
    z-index: 2;
  }
  .Initiate_mask2{
    position: relative;
    margin-top: -45%;
    width: 100%;
    z-index: 2;
  }
  .Initiate_mask3{
    position: fixed;
    top: 20%;
    width: 100%;
    z-index: 2;
  }
  .Initiate_mask4{
    position: fixed;
    top: 100px;
    width: 100%;
    z-index: 2;
  }
  .Initiate_crowdfunding{
    margin: 0 auto;
    width: 950px;
    height: 500px;
  }
  .select_title{
    width: 100%;
    height: 50px;
    display: flex;
    margin-bottom: 90px;
  }
  .select_title p{
    width: calc(100% - 26px);
    text-align: center;
    font-size: 32px;
  }
  .select_title img{
    width: 26px;
    height: 26px;
    position: relative;
    top: 12px;
    cursor: pointer;
  }
  .window{
    height: 320px;
  }
  .window_img_yl{
    margin-top: 70px;
  }
  .window img{
    margin-top: 90px;
  }
  .window p{
    color: #0071FE;
    cursor: pointer;
  }
  .window p:nth-child(2){
    font-size: 24px;
    font-weight: bold;
    margin-top: 10px;
  }
  .window p:nth-child(3){
    font-size: 16px;
  }
  /* 弹框 */
  .window_tk{
    position: relative;
    margin: 0 auto;
    width: 1110px;
    height: auto;
    background: #fff;
    border-radius: 20px;
    z-index: 2;
  }
  .window_tk2{
    position: relative;
    margin: 0 auto;
    width: 700px;
    height: auto;
    background: #fff;
    border-radius: 20px;
    box-shadow: 0px 0px 4px #ccc;
    z-index: 2;
  }
  .window_tk_title{
    border-top-left-radius: 20px;
    border-top-right-radius: 20px;
    width: 100%;
    height: 173px;
    display: flex;
    background-image: linear-gradient(to bottom, #2987FE, #FFFFFF);
    
  }
  .window_tk_title p{
    width: calc(100% - 86px);
    text-align: left;
    font-size: 22px;
    margin-left: 30px;
    line-height: 50px;
  }
  .window_tk_title img{
    width: 26px;
    height: 26px;
    position: relative;
    top: 12px;
    cursor: pointer;
  }
  .window_tk_content{
    position: relative;
    top: -120px;
    min-height: 300px;
    border-radius: 20px;
    width: 1056px;
    margin: 0 auto;
    background: #fff;
  }
  .window_tk_content2{
    position: relative;
    top: -120px;
    min-height: 200px;
    border-radius: 20px;
    width: 656px;
    margin: 0 auto;
    background: #fff;
  }
  .window_btn{
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width:240px;
    height: 60px;
    line-height: 60px;
    font-size: 23px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #0071FE, #0071FE);
  }
  .window_btn2{
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width:240px;
    height: 60px;
    line-height: 60px;
    font-size: 23px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .window_jk_title{
    position: relative;
    top: -40px;
    margin: 0 auto;
    width:240px;
    height: 60px;
    line-height: 60px;
    font-size: 16px;
    color: red;
    text-align: center;
    border-radius: 50px;
  }
  .window_tk_nr{
    width: 92%;
    margin: 0 auto;
    padding-top: 20px;
  }
  .window_tk_nr_width{
    width: 300px;
  }
  .window_tk_nr_p{
    color: #000;
    padding-bottom: 10px;
    position: relative;
    left: -20px;
  }
  .img_text{
    position: relative;
    color: #aaa;
    left: 140px;
    top: -30px;
    font-size: 12px;
  }
  .details{
    margin-left: 10px;
  }
  .details>p{
    color: #000;
    font-size: 14px;
    padding: 10px;
    position: relative;
    left:-20px;
  }
  .details>p span{
    color: #0071FE;
  }
  .window_xx_name{
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
    display: flex;
  }
  .window_tx{
    width: 46px;
    height: 46px;
    border-radius: 25px;
  }
  .window_xx_name p{
    height: 46px;
    line-height: 46px;
    font-size: 22px;
    margin-left: 20px;
  }
  .window_xx_name p:nth-child(3){
    font-size: 18px!important;
    color:#A1A1A1;
  }
  .window_xx_title{
    color: #000;
    width: 95%;
    margin: 0 auto;
    margin-top: 23px;
    height: 30px;
    line-height: 30px;
    display: flex;
  }
  .window_xx_title p:nth-child(1){
    width: calc(100% - 220px);
    font-size: 28px;
    font-weight: bold;
    overflow: hidden;
   text-overflow:ellipsis;
   white-space: nowrap;
  }
  .window_xx_title p:nth-child(2){
    width: 220px;
    font-size: 18px;
    color: #a1a1a1;
  }
  .window_xx_content{
    width: 95%;
    margin: 0 auto;
    margin-top: 50px;
  }
  .window_xx_content p{
    text-indent: 35px;
    line-height: 30px;
    font-size: 18px;
    color: #373737;
  }
  .window_xx_img img{
    margin-top: 20px;
    width: 200px;
    height: 200px;
  }
  .window_xx_name2{
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
  }
  .window_tx_left{
    float:left;
    display: flex;
  }
  .window_tx_left p{
    height: 46px;
    line-height: 46px;
    font-size: 22px;
    margin-left: 20px;
  }
  .window_tx_left p:nth-child(3){
    font-size: 18px!important;
    color:#A1A1A1;
  }
  .window_tx_right{
    width: 165px;
    height: 44px;
    line-height: 44px;
    text-align: center;
    background: #F65F47;
    border-radius: 50px;
    color: #fff;
    font-size: 22px;
    float: right;
  }
  .window_xs_content{
    width: 95%;
    margin: 0 auto;
    margin-top: 20px;
    background: #F8F8F8;
    border-radius: 10px;
    height: auto;
  }
  .window_col{
    text-align: center;
    color: #000;
    margin-top: 10px;
  }
  .window_col p:nth-child(1){
    color: #FE6700;
    font-size: 32px;
  }
  .window_col p:nth-child(2){
    color: #707070;
    font-size: 20px;
  }
  .window_col2{
    text-align: center;
    height: 120px;
    line-height: 120px;
    color: #000;
    margin-top: 60px;
    background: #fff;
    box-shadow: 0px 0px 4px #ccc;
    border-radius: 15px;
    font-size: 20px;
    border:2px solid #fff;
    cursor: pointer;
  }
  .window_col2:hover{
    border:2px solid #0071FE;
  }
  .window_col3{
    border:2px solid #0071FE;
  }
  .window_xs_lable{
    width: 95%;
    margin: 0 auto;
    margin-top: 20px;
    background: #fff;
    border-radius: 10px;
    height: auto;
  }
  .window_xs_lable_title{
    width: 95%;
    margin: 0 auto;
    padding: 20px 0;
  }
  .window_xs_lable_title .left{
    float: left;
    color: #5E5E5E;
    font-size: 22px;
  }
  .window_xs_lable_title .right{
    float: right;
    color: #A1A1A1;
    font-size: 16px;
  }
  .window_xs_lable_content{
    width: 95%;
    margin: 0 auto;
    color: #000;
  }
  .window_xs_lable_content>div{
    width: 100%;
    display: flex;
    height: 50px;
    line-height: 50px;
    font-size: 20px;
  }
  .window_xs_lable_content>div>p{
    width: 33.33%;
    text-align: left;
  }
  .window_xs_lable_content>div>p:nth-child(2){
    color: #F16219;
    text-align: center;
  }
  .window_xs_lable_content>div>p:nth-child(3){
    text-align: right;
  }
  .window_xs_pl{
    width: 100%;
    display: flex;
  }
  .window_xs_pl p{
    width: 50%;
  }
  .window_xs_pl p:nth-child(1){
    text-align: left;
    font-size: 20px;
    line-height: 30px;
  }
  .window_xs_pl_img{
    width: 30px;
    height: 30px;
    border-radius: 50px;
    position: relative;
    top: 8px;
  }
  .window_xs_pl p:nth-child(2){
    text-align: right;
    font-size:18px;
    position: relative;
    top: 8px;
  }
  .window_xs_pl_con{
    width:100%;
    line-height: 30px;
    border-bottom: 1px solid #ccc;
    padding:20px 0;
    text-indent: 20px;
  }
  .window_jkjg{
    width: 500px;
    margin: 0 auto;
  }
  .swiper-container{
    min-height: 300px;
  }
  /*自适应手机*/
@media screen and (max-width:1200px){
  .swiper-container{
    min-height: 250px;
  }
  ::v-deep.tabs1{
      .ivu-tabs-bar{
        .ivu-tabs-nav-container{
          .ivu-tabs-nav-wrap{
            .ivu-tabs-nav-scroll{
              .ivu-tabs-nav{
                .ivu-tabs-tab{
                  font-size: 22px;
                  width: 50%;
                  text-align: center;
                  margin-left: 5px;
                }
                .ivu-icon{
                  width: 24px;
                  height: 24px;
                }
                .ivu-tabs-ink-bar{
                  width: 135px!important;
                  margin: 0 auto;
                  margin-left: 15px;
                }
              }
            }
          }
        }
      }
    }
  ::v-deep.tabs2{
      .ivu-tabs-bar{
        .ivu-tabs-nav-container{
          .ivu-tabs-nav-wrap{
            .ivu-tabs-nav-scroll{
              .ivu-tabs-nav{
                width: 100%;
                .ivu-tabs-tab{
                  font-size: 16px;
                  width: 50%;
                  text-align: center;
                  margin-left: 0px;
                }
                .ivu-icon{
                  width: 18px;
                  height: 18px;
                }
                .ivu-tabs-ink-bar{
                  width: 175px!important;
                }
              }
            }
          }
        }
      }
    }
  ::v-deep.yanse2{
      position: relative;
      top: 30px;
      .ivu-input{
        color: #FD4E17!important;
        background-color: #fff;
        color: #000;
        height: 40px;
        border-color: #f5f5f5;
        box-shadow: 0px 0px 4px #ccc;
        border-radius: 10px;
      }
    }
  .activity {
    width: 100%;
    background: rgba(242,246,250,1) !important;
    height: 100%;
    background-size: cover;
    position: relative;
    overflow: hidden;
    padding-bottom: 20px;
    padding-top: 45px;
    color: #fff;
  }
  .activity .banner{
    width:100%;
    height: 400px;
  }
  .banner_text{
    position: absolute;
    top: 60px;
    width: 100%;
    text-align: center;
    color: #fff;
    font-weight: 500;
    font-size: 24px;
    margin-top: 0px;
  }
  .br{
    display: block;
  }
  .money{
    font-size: 36px;
    color: red;
    letter-spacing:3px;
  }
  .money_span{
    font-size: 24px;
    color: #fff;
    font-weight: 400;
    position: relative;
    top: -0px;
  }
  .banner_content{
    width: 100%;
    margin: 0 auto;
    margin-top: 10px;
  }
  .banner_content p{
    font-size: 20px;
    letter-spacing:1px;
    width: 50%;
    text-align: center;
  }
  .banner_content p span{
    font-size: 18px;
    letter-spacing:1px;
    position: relative;
    top:-1px;
  }
  .banner_content p:nth-child(1){
    float: left;
    font-weight: 600;
  }
  .banner_content p:nth-child(2){
    float: right;
    font-weight: 600;
  }
  .tab{
    width: 100%;
    font-weight: 500;
  }
  .tab_all{
    display: none;
  }
  .swiper-wrapper{
    margin: 0px 0px;
  }
  .swiper-button-prev,
  .swiper-button-next{
    display: none;
  }
  .swiper-slide2 img {
    	width: 100%;
      border-top-left-radius: 14px;
      border-top-right-radius: 14px;
    	display: block;
      height: 224px;
  }
  .swiper-slide2>div {
    	width: auto;
      box-shadow: 0px 0px 4px #ccc;
      border-radius: 14px;
      height: 300px;
  }
  .swiper-slide2 .swiper_title{
      width: 95%;
      display: flex;
      margin-top:25px;
  }
  .btn{
    margin: 0 auto;
    margin-top: 20px;
    width:190px;
    height: 45px;
    line-height: 45px;
    font-size: 18px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .btn span{
    cursor: pointer;
  }
  .bottom{
    width: 100%;
    text-align: center;
    margin-top: 42px;
  }
  .bottom img{
    width: 95%;
    margin: 0 auto;
  }
  .bottom_text{
    width: 95%;
    margin: 0 auto;
    text-align: left;
    color: #2987FE;
  }
  .tab_all2{
    width: 100%;
    text-align: center;
    font-size: 18px;
    margin-top: 10px;
    display: block;
  }
  .Initiate_mask4{
    position: fixed;
    top: 40px;
    width: 100%;
    z-index: 2;
  }
  .all_xs_img{
      float: left;
      width: 100%;
    }
    .all_xs_img img{
      width: 100%;
      height: 160px;
      border-radius: 10px;
    }
  .all_xs_text{
      float: none;
      width: auto;
      color: #000;
      border-radius: 10px;
    }
    .all_xs_text_left{
      float: none;
      margin-left: 0px;
    }
    .all_xs_text_title{
      width: 300px;
      font-size: 20px;
      overflow:hidden;
      text-overflow:ellipsis;
      white-space:nowrap;
    }
    .all_xs_text_content{
      background: #f5f5f5;
      border-radius: 10px;
      padding: 0px 0px;
      display: flex;
    }
    .all_xs_text_content>div{
      width: 50%;
      text-align: center;
    }
    .all_xs_text_content>div:nth-child(2){
      display: none;
    }
    .all_xs_text_content p{
      line-height: 30px;
      min-width: 100px;
      padding: 5px 10px;
      text-align: center;
    }
    .all_xs_text_content p:nth-child(2){
      color: red;
      font-size: 20px;
      font-weight: bold;
    }
    .all_xs_text_content p:nth-child(1){
      color: #000;
      font-size: 16px;
      white-space:nowrap;
    }
    .all_xs_text_money{
      position: relative;
      top: 5px;
      color: #FE6700;
    }
    .all_xs_text_right{
      float: none;
      position: relative;
    }
    .all_xs_right_top{
      font-size: 16px;
      height: 40px;
      line-height: 40px;
      float: left;
    }
    .all_xs_right_top span:nth-child(1){
      color: red;
      font-size: 20px;
    }
    .all_xs_right_top span:nth-child(2){
      color: #666;
      font-size: 14px;
      position: relative;
      top:0px;
    }
    .all_xs_right_bottom{
      float: right;
      margin-top: 12px;
      font-size: 14px;
    }
  .Initiate_mask{
    position: fixed;
    top: 30%;
    width: 100%;
    z-index: 2;
  }
  .Initiate_crowdfunding{
    margin: 0 auto;
    width: 95%;
    height: 300px;
  }
  .select_title{
    width: 100%;
    height: 50px;
    display: flex;
    margin-bottom: 10px;
  }
  .select_title p{
    width: calc(100% - 26px);
    text-align: center;
    font-size: 22px;
  }
  .select_title img{
    width: 22px;
    height: 22px;
    position: relative;
    top: 6px;
    cursor: pointer;
  }
  .window{
    height: 120px;
  }
  .window img{
    margin-top: 10px;
    width: 30px;
    height: 30px;
  }
  .window_img_yl{
    margin-top: 10px;
  }
  .window p{
    color: #0071FE;
    cursor: pointer;
  }
  .window p:nth-child(2){
    font-size: 14px;
    font-weight: bold;
    margin-top: 10px;
  }
  .window p:nth-child(3){
    font-size: 12px;
  }
   /* 弹框 */
  .Initiate_mask2{
    position: relative;
    margin-top: -840px;
    width: 100%;
    z-index: 2;
  }
  .window_tk{
    position: relative;
    margin: 0 auto;
    width: 100%;
    height: auto;
    background: #fff;
    border-radius: 0px;
    z-index: 2;
  }
  .window_tk_title{
    border-top-left-radius: 0px;
    border-top-right-radius: 0px;
    width: 100%;
    height: 173px;
    display: flex;
    background-image: linear-gradient(to bottom, #2987FE, #FFFFFF);
    
  }
  .window_tk_title p{
    width: calc(100% - 46px);
    text-align: left;
    font-size: 18px;
    margin-left: 10px;
    line-height: 50px;
  }
  .window_tk_title img{
    width: 18px;
    height: 18px;
    position: relative;
    top: 16px;
    cursor: pointer;
  }
  .window_tk_content{
    position: relative;
    top: -120px;
    min-height: 300px;
    border-radius: 0px;
    width: 100%;
    margin: 0 auto;
    background: #fff;
  }
  .window_btn{
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width:140px;
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #0071FE, #0071FE);
  }
  .window_btn2{
    cursor: pointer;
    position: relative;
    top: -40px;
    margin: 0 auto;
    width:140px;
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    color: #fff;
    text-align: center;
    border-radius: 50px;
    background-image: linear-gradient(to bottom, #FF9113, #F16219);
  }
  .window_tk_nr{
    width: 92%;
    margin: 0 auto;
    padding-top: 20px;
  }
  .window_tk_nr_width{
    width: 300px;
  }
  .window_tk_nr_p{
    color: #000;
    padding-bottom: 10px;
    position: relative;
    left: -5px;
  }
  .img_text{
    position: relative;
    color: #aaa;
    left: 0px;
    top: 8px;
    font-size: 12px;
  }
  .details{
    margin-left: 10px;
  }
  .details>p{
    color: #000;
    font-size: 14px;
    padding: 10px;
    position: relative;
    left:-20px;
  }
  .details>p span{
    color: #0071FE;
  }
  .window_xx_name{
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
    display: flex;
  }
  .window_tx{
    width: 40px;
    height: 40px;
    border-radius: 25px;
  }
  .window_xx_name p{
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    margin-left: 20px;
  }
  .window_xx_name p:nth-child(3){
    font-size: 14px!important;
    color:#A1A1A1;
  }
  .window_xx_title{
    color: #000;
    width: 95%;
    margin: 0 auto;
    margin: 10px 15px;
    height: 45px;
    line-height: 30px;
    display: block;
  }
  .window_xx_title p:nth-child(1){
    width: 100%;
    font-size: 16px;
    font-weight: bold;
  }
  .window_xx_title p:nth-child(2){
    width: 200px;
    font-size: 14px;
    color: #a1a1a1;
  }
  .window_xx_content{
    width: 95%;
    margin: 0 auto;
    margin-top: 50px;
  }
  .window_xx_content p{
    text-indent: 35px;
    line-height: 30px;
    font-size: 14px;
    color: #373737;
  }
  .window_xx_img img{
    margin-top: 20px;
    width: 150px;
    height: 150px;
  }
  .window_xx_name2{
    width: 95%;
    margin: 0 auto;
    padding-top: 20px;
    color: #000;
  }
  .window_tx_left{
    float:left;
    display: flex;
  }
  .window_tx_left p{
    height: 40px;
    line-height: 40px;
    font-size: 16px;
    margin-left: 10px;
  }
  .window_tx_left p:nth-child(3){
    font-size: 14px!important;
    color:#A1A1A1;
  }
  .window_tx_right{
    width: 105px;
    margin-top: 3px;
    height: 34px;
    line-height: 34px;
    text-align: center;
    background: #F65F47;
    border-radius: 50px;
    color: #fff;
    font-size: 14px;
    float: right;
  }
  .window_xs_content{
    width: 95%;
    margin: 0 auto;
    margin-top: 20px;
    background: #F8F8F8;
    border-radius: 10px;
    height: auto;
  }
  .window_col{
    text-align: center;
    color: #000;
    margin-top: 10px;
  }
  .window_col p:nth-child(1){
    color: #FE6700;
    font-size: 16px;
  }
  .window_col p:nth-child(2){
    color: #707070;
    font-size: 12px;
  }
  .window_xs_lable{
    width: 95%;
    margin: 0 auto;
    margin-top: 12px;
    background: #fff;
    border-radius: 10px;
    height: auto;
  }
  .window_xs_lable_title{
    width: 95%;
    margin: 0 auto;
    padding: 10px 0;
  }
  .window_xs_lable_title .left{
    float: left;
    color: #5E5E5E;
    font-size: 14px;
  }
  .window_xs_lable_title .right{
    float: right;
    color: #A1A1A1;
    font-size: 12px;
    margin-top: 2px;
  }
  .window_xs_lable_content{
    width: 95%;
    margin: 0 auto;
    color: #000;
  }
  .window_xs_lable_content>div{
    width: 100%;
    display: flex;
    height: 40px;
    line-height: 40px;
    font-size: 14px;
  }
  .window_xs_lable_content>div>p{
    width: 33.33%;
    text-align: left;
  }
  .window_xs_lable_content>div>p:nth-child(1){
    width: 40%;
    text-align: left;
  }
  .window_xs_lable_content>div>p:nth-child(2){
     width: 30%;
    color: #F16219;
    text-align: center;
  }
  .window_xs_lable_content>div>p:nth-child(3){
     width: 30%;
    text-align: right;
  }
  .window_xs_pl{
    width: 100%;
    display: flex;
  }
  .window_xs_pl p{
    width: 50%;
  }
  .window_xs_pl p:nth-child(1){
    text-align: left;
    font-size: 14px;
    line-height: 25px;
  }
  .window_xs_pl_img{
    width: 25px;
    height: 25px;
    border-radius: 50px;
    position: relative;
    top: 8px;
  }
  .window_xs_pl p:nth-child(2){
    text-align: right;
    font-size:14px; 
    line-height: 25px;
  }
  .window_xs_pl_con{
    width:100%;
    line-height: 30px;
    border-bottom: 1px solid #ccc;
    padding:10px 0;
    text-indent: 20px;
  }
  .Initiate_mask3{
    position: fixed;
    top: 15%;
    width: 100%;
    z-index: 2;
  }
  .window_tk2{
    position: relative;
    margin: 0 auto;
    width: 95%;
    height: auto;
    background: #fff;
    border-radius: 0px;
    box-shadow: 0px 0px 4px #ccc;
    z-index: 2;
  }
  .window_tk_content2{
    position: relative;
    top: -120px;
    min-height: 200px;
    border-radius: 20px;
    width: 90%;
    margin: 0 auto;
    background: #fff;
  }
  .window_col2{
    text-align: center;
    height: 100px;
    line-height: 100px;
    color: #000;
    margin-top: 60px;
    background: #fff;
    box-shadow: 0px 0px 4px #ccc;
    border-radius: 15px;
    font-size: 14px;
    border:2px solid #fff;
    cursor: pointer;
  }
  .window_col2:hover{
    border:2px solid #0071FE;
  }
  .window_jkjg{
    width: 80%;
    margin: 0 auto;
  }
  .window_jkjg{
    width: 95%;
    margin: 0 auto;
  }
}
@media screen and (min-width:1300px){
  .Initiate_mask2{
    margin-top: -70%;
  }
}
@media screen and (min-width:1300px){
  .Initiate_mask2{
    margin-top: -70%;
  }
}
@media screen and (min-width:1400px){
  .Initiate_mask2{
    margin-top: -65%;
  }
}
@media screen and (min-width:1500px){
  .Initiate_mask2{
    margin-top: -60%;
  }
}
@media screen and (min-width:1600px){
  .Initiate_mask2{
    margin-top: -55%;
  }
}
@media screen and (min-width:1600px){
  .Initiate_mask2{
    margin-top: -50%;
  }
}
@media screen and (min-width:1900px){
  .Initiate_mask2{
    margin-top: -45%;
  }
}
</style>
