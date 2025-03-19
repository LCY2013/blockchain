

<style lang="scss" scoped>
.entrustcurrent {
  float: left;
  width: 100%;
  padding-left: 30px;
}

.page {
  text-align: right;
  margin-top: 20px;
}
.table {
  border-radius: 4px;
}
.table .ivu-table-wrapper {
  position: relative;
  /* border: 1px solid #dddee1; */
  border-bottom: 0;
  border-right: 0;
  // box-shadow: 0 0 2px #ccc;
  border-radius: 4px;
  // overflow: hidden;
}
.form.ivu-form-inline .ivu-form-item {
  display: inline-block;
}
</style>
<style lang="scss">
.entrustcurrent .ivu-table th,
.entrustcurrent .ivu-table td {
  text-align: center;
}
.table .ivu-table-cell-expand {
  color: #f0a70a;
}
.order-type-btn-wrap {
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  -ms-flex-pack: justify;
  justify-content: space-between;
  height: 30px;
  padding: 0 16px;
  min-width: 1320px;
  background:#192330;
  button {
    width: 80px;
    border-radius:5px;
    height:auto;
    background: #192330;
    border: 1px solid #fff;
    padding:3px 5px;
    text-align: center;
    font-size: 12px;
    margin-right: 10px;
    color: #fff;
  }
  .btn-selected {
    color: #000 !important;
    background: #f0ac19;
    border: 1px solid #f0ac19 !important;
  }
}
.buy {
  color: #00b275 !important;
}
.sell {
  color: #f15057 !important;
}
.ivu-tabs-nav .ivu-tabs-tab-active {
    color: #f0a70a;
}
.ivu-tabs-nav .ivu-tabs-tab:hover {
    color: #f0a70a;
}
.ivu-tabs-ink-bar {
  background-color: #f0a70a;
}
</style>

<template>
  <div class="entrustcurrent">
    <div class="table">
      <Tabs :active-key="formItem.type" @on-click="setSelectedType" :animated="false">
        <Tab-pane :label="$t('uc.myCf.medicalCrowdfunding')" :key="1"></Tab-pane>
        <Tab-pane :label="$t('uc.myCf.otherCrowdfunding')" :key="2"></Tab-pane>
        <Tab-pane :label="$t('uc.myCf.offlineCharity')" :key="3"></Tab-pane>
        <Tab-pane :label="$t('uc.myCf.myDonation')" :key="4"></Tab-pane>
      </Tabs>
      <Table v-show="formItem.type==1" :no-data-text="$t('common.nodata')" :columns="columns" :data="ylList" :loading="loading"></Table>
      <Table v-show="formItem.type==2" :no-data-text="$t('common.nodata')" :columns="columns2" :data="qtList" :loading="loading"></Table>
      <Table v-show="formItem.type==3" :no-data-text="$t('common.nodata')" :columns="columns3" :data="gyList" :loading="loading"></Table>
      <Table v-show="formItem.type==4" :no-data-text="$t('common.nodata')" :columns="columns4" :data="jkList" :loading="loading"></Table>
      <div class="page">
        <Page  v-show="formItem.type==1" :total="total" :pageSize="pageSize" :current="pageNo" @on-change="loadDataPage"></Page>
        <Page  v-show="formItem.type==2" :total="total2" :pageSize="pageSize2" :current="pageNo" @on-change="loadDataPage2"></Page>
        <Page  v-show="formItem.type==3" :total="total3" :pageSize="pageSize3" :current="pageNo" @on-change="loadDataPage3"></Page>
        <Page  v-show="formItem.type==4" :total="total4" :pageSize="pageSize4" :current="pageNo" @on-change="loadDataPage4"></Page>
      </div>
    </div>
  </div>
</template>
<script>
var moment = require("moment");
import expandRow from "@components/exchange/expand.vue";

export default {
  components: { expandRow },
  data() {
    const self = this;
    return {
      loading: false,
      pageSize: 10,
      pageNo: 1,
      total: 10,
      pageSize2: 10,
      pageNo2: 1,
      total2: 10,
      pageSize3: 10,
      pageNo3: 1,
      total3: 10,
      pageSize4: 10,
      pageNo4: 1,
      total4: 10,
      symbol: [],
      formItem: {
        symbol: "",
        type: 1,
        entrustType: "",
        direction: "",
        date: ""
      },
      columns: [
        {
          title: self.$t("uc.myCf.title"),
          key: "fundingTitle",
          render(h, params) {
            return h("span",params.row.fundingTitle);
          },
        },
        {
            title: self.$t("uc.myCf.time"),
            key: "addTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.addTime));
            },
        },
        {
          title: self.$t("uc.myCf.targetAmount"),
          key: "targetAmount",
          render(h, params) {
            return h(
              "span",
              params.row.targetAmount + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.amountRaised"),
          key: "amountReceived",
          render(h, params) {
            return h(
              "span",
              params.row.amountReceived + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.withdrawable"),
          key: "drawMoney",
          render(h, params) {
            return h(
              "span",
              params.row.drawMoney + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.status"),
          key: "status",
          render(h, params) {
             const status = params.row.status;
             if (status == 0) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#F3BB01",
                    },
                  },
                  self.$t("uc.myCf.pending")
                );
              } else if (status == 1) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#05DF3F",
                    },
                  },
                  self.$t("uc.myCf.passed")
                );
              } else if (status == 2) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#E80000",
                    },
                  },
                  self.$t("uc.myCf.rejected")
                );
              } else {
                return h("span", 
                {
                    style: {
                      color: "#A7A5A1",
                    },
                }, 
                self.$t("uc.myCf.toEnd")
                );
              }
          },
        },
                {
            title: self.$t("uc.myCf.operating"),
            width: 110,
            render: (h, params) => {
              if(params.row.drawMoney != 0){
              return h(
                "Button",
                {
                  props: {
                    size: "small",
                  },
                  style: {
                    border: "1px solid #f0ac19",
                    color: "#f1ac19",
                    "line-height": "1.2",
                    "border-radius": "10px",
                  },
                  on: {
                    click: () => {
                      // console.log("======开始撤单")
                      this.withdraw(params.row);
                    },
                  },
                },
                self.$t("uc.myCf.withdraw")
              );
              }else{
                return h(
                "Button",
                {
                  props: {
                    size: "small",
                  },
                  style: {
                    border: "1px solid #f0ac19",
                    color: "#f1ac19",
                    disabled,
                    "line-height": "1.2",
                    "border-radius": "10px",
                  },
                },
                self.$t("uc.myCf.withdraw")
              );
              };
            },
          },
      ],
      columns2: [
        {
          title: self.$t("uc.myCf.title"),
          key: "fundingTitle",
          render(h, params) {
            return h("span",params.row.fundingTitle);
          },
        },
        {
            title: self.$t("uc.myCf.time"),
            key: "addTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.addTime));
            },
        },
        {
            title: self.$t("uc.myCf.type"),
            key: "type",
            render: (h, params) => {
              return h(
                "span", 
                params.row.type === 1 ? "心愿" : "创业"
              );
            },
        },
        {
          title: self.$t("uc.myCf.targetAmount"),
          key: "targetAmount",
          render(h, params) {
            return h(
              "span",
              params.row.targetAmount + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.amountRaised"),
          key: "amountReceived",
          render(h, params) {
            return h(
              "span",
              params.row.amountReceived + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.withdrawable"),
          key: "drawMoney",
          render(h, params) {
            return h(
              "span",
              params.row.drawMoney + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.status"),
          key: "status",
          render(h, params) {
             const status = params.row.status;
             if (status == 0) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#F3BB01",
                    },
                  },
                  self.$t("uc.myCf.pending")
                );
              } else if (status == 1) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#05DF3F",
                    },
                  },
                  self.$t("uc.myCf.passed")
                );
              } else if (status == 2) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#E80000",
                    },
                  },
                  self.$t("uc.myCf.rejected")
                );
              } else {
                return h("span", 
                {
                    style: {
                      color: "#A7A5A1",
                    },
                }, 
                self.$t("uc.myCf.toEnd")
                );
              }
          },
        },
        {
            title: self.$t("uc.myCf.operating"),
            width: 110,
            render: (h, params) => {
              if(params.row.drawMoney != 0){
              return h(
                "Button",
                {
                  props: {
                    size: "small",
                  },
                  style: {
                    border: "1px solid #f0ac19",
                    color: "#f1ac19",
                    "line-height": "1.2",
                    "border-radius": "10px",
                  },
                  on: {
                    click: () => {
                      // console.log("======开始撤单")
                      this.withdraw(params.row);
                    },
                  },
                },
                self.$t("uc.myCf.withdraw")
              );
              }else{
                return h(
                "Button",
                {
                  props: {
                    size: "small",
                  },
                  style: {
                    border: "1px solid #f0ac19",
                    color: "#f1ac19",
                    disabled,
                    "line-height": "1.2",
                    "border-radius": "10px",
                  },
                },
                self.$t("uc.myCf.withdraw")
              );
              };
            },
          },
      ],
      columns3: [
        {
          title: self.$t("uc.myCf.title"),
          key: "fundingTitle",
          render(h, params) {
            return h("span",params.row.fundingTitle);
          },
        },
        {
            title: self.$t("uc.myCf.time"),
            key: "addTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.addTime));
            },
        },
        {
            title: self.$t("uc.myCf.type"),
            key: "type",
            render: (h, params) => {
              const type = params.row.type;
              if (type == 1) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#F65F47",
                    },
                  },
                  self.$t("uc.myCf.wish")
                );
              } else if (type == 2) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#0071FE",
                    },
                  },
                  self.$t("uc.myCf.pioneer")
                );
              }
            },
        },
        {
          title: self.$t("uc.myCf.applyAmount"),
          key: "applyAmount",
          render(h, params) {
            return h(
              "span",
              params.row.applyAmount + ' USDT'
            );
          },
        },
        {
          title: self.$t("uc.myCf.status"),
          key: "status",
          render(h, params) {
             const status = params.row.status;
             if (status == 0) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#F3BB01",
                    },
                  },
                  self.$t("uc.myCf.pending")
                );
              } else if (status == 1) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#05DF3F",
                    },
                  },
                  self.$t("uc.myCf.passed")
                );
              } else if (status == 2) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#E80000",
                    },
                  },
                  self.$t("uc.myCf.rejected")
                );
              } else {
                return h("span", 
                {
                    style: {
                      color: "#A7A5A1",
                    },
                }, 
                self.$t("uc.myCf.toEnd")
                );
              }
          },
        },
      ],
      columns4: [
        {
          title: self.$t("uc.myCf.title"),
          key: "fundingTitle",
          render(h, params) {
            return h("span",params.row.fundingTitle);
          },
        },
        {
            title: self.$t("uc.myCf.time"),
            key: "fundingTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.fundingTime));
            },
        },
        {
            title: self.$t("uc.myCf.type"),
            key: "fundingType",
            render: (h, params) => {
              const type = params.row.fundingType;
              if (type == 1) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#F65F47",
                    },
                  },
                  self.$t("uc.myCf.wish")
                );
              } else if (type == 2) {
                return h(
                  "span",
                  {
                    style: {
                      color: "#0071FE",
                    },
                  },
                  self.$t("uc.myCf.pioneer")
                );
              } else {
                return h(
                  "span", 
                  {
                      style: {
                        color: "#0AC159",
                      },
                    },
                    self.$t("uc.myCf.medical")
                );
              }
            },
          },
        {
          title: self.$t("uc.myCf.donateMoney"),
          key: "fundingMoney",
          render(h, params) {
            return h(
              "span",
              params.row.fundingMoney + ' USDT'
            );
          },
        }
      ],
      ylList: [],
      qtList: [],
      gyList: [],
      jkList: [],
    };
  },
  computed:{
    lang: function() {
      return this.$store.state.lang;
    }
  },
  watch: {
    lang: function() {
      this.updateLangData();
    }
  },
  created() {
    this.refresh();
    this.refresh2();
    this.refresh3();
    this.refresh4();
  },
  methods: {
    withdraw(e){
      console.log(e)
      let param = {}
          param.fundingId = e.id //目标金额
          param.type = e.type
          console.log(param)
      this.$http
        .post(this.host + this.api.crowdfunding.getCash,param)
        .then(response => {
          var resp = response.body;
          if (resp.code == 0) {
            this.$Message.success(resp.message);
            this.refresh1();
            this.refresh2();
          } else {
            this.$Message.error(resp.message);
          }
        });
    },
    setSelectedType(e) {
      console.log(e)
      this.handleClear();
      if (e==0) { //第一个
        this.formItem.type = 1; //医疗众筹
      } else if (e==1) {
        this.formItem.type = 2; //其他众筹
      } else if (e==2) {
        this.formItem.type = 3; //线下公益
      } else if (e==3) {
        this.formItem.type = 4; //我的捐款
      }
      this.pageNo = 1;
      this.refresh();
      this.refresh2();
      this.refresh3();
      this.refresh4();
    },
    dateFormat: function(tick) {
      return moment(tick).format("YYYY-MM-DD");
    },
    timeFormat: function(tick) {
      return moment(tick).format("HH:mm:ss");
    },
    loadDataPage(data) {
      this.pageNo = data;
      this.refresh();
    },
    loadDataPage2(data) {
      this.pageNo2 = data;
      this.refresh2();
    },
    loadDataPage3(data) {
      this.pageNo3 = data;
      this.refresh3();
    },
    loadDataPage4(data) {
      this.pageNo4 = data;
      this.refresh4();
    },
    handleClear() {
      this.formItem = {
        symbol: "",
        direction: "",
        date: "",
        entrustType: "",
        type: this.formItem.type
      };
    },
    refresh() {
      //查询历史委托
      this.loading = true;
      let params = {};
      params.pageNo = this.pageNo-1;
      params.pageSize = this.pageSize;
      var that = this;
      this.ylList = [];
      this.$http
        .post(this.host + this.api.crowdfunding.myCrowdfundingYl,params,{emulateJSON:false})
        .then(response => {
          var resp = response.body.data;
          console.log(resp);
          if (resp.content.length > 0) {
            this.total = resp.totalElements;
            this.ylList = resp.content;
          }
          this.loading = false;
        });
    },
    refresh2() {
      //查询历史委托
      this.loading = true;
      let params = {};
      params.pageNo = this.pageNo2-1;
      params.pageSize = this.pageSize2;
      this.qtList = [];
      this.$http
        .post(this.host + this.api.crowdfunding.myCrowdfundingQt,params,{emulateJSON:false})
        .then(response => {
          var resp = response.body.data;
          console.log(resp);
          if (resp.content.length > 0) {
            this.total2 = resp.totalElements;
            this.qtList = resp.content;
          }
          this.loading = false;
        });
    },
    refresh3() {
      //查询历史委托
      this.loading = true;
      let params = {};
      params.pageNo = this.pageNo3-1;
      params.pageSize = this.pageSize3;
      this.gyList = [];
      this.$http
        .post(this.host + this.api.crowdfunding.myCrowdfundingGy,params,{emulateJSON:false})
        .then(response => {
          var resp = response.body.data;
          console.log(resp);
          if (resp.content.length > 0) {
            this.total3 = resp.totalElements;
            this.gyList = resp.content;
          }
          this.loading = false;
        });
    },
    refresh4() {
      //查询历史委托
      this.loading = true;
      let params = {};
      params.pageNo = this.pageNo4-1;
      params.pageSize = this.pageSize4;
      this.jkList = [];
      this.$http
        .post(this.host + this.api.crowdfunding.myCrowdfundingJk,params,{emulateJSON:false})
        .then(response => {
          var resp = response.body.data;
          console.log(resp);
          if (resp.content.length > 0) {
            this.total4 = resp.totalElements;
            this.jkList = resp.content;
          }
          this.loading = false;
        });
    },
    updateLangData(){
      this.columns[1].title = this.$t("exchange.time");
      this.columns[2].title = this.$t("uc.finance.trade.symbol");
      this.columns[3].title = this.$t("uc.finance.trade.type");
      this.columns[4].title = this.$t("exchange.direction");
      this.columns[5].title = this.$t("exchange.price");
      this.columns[6].title = this.$t("exchange.num");
      this.columns[7].title = this.$t("exchange.traded");
      this.columns[8].title = this.$t("uc.finance.trade.turnover");
      this.columns[9].title = this.$t("exchange.action");
    }
  }
};
</script>

