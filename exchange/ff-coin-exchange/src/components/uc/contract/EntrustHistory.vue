
<style lang="scss" scoped>
.entrusthistory {
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
.entrusthistory .ivu-table th,
.entrusthistory .ivu-table td {
  text-align: center;
}
.table .ivu-table-cell-expand {
  color: #f0a70a;
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
  <div class="entrusthistory">
    <Form class="form" :model="formItem" :label-width="75" inline>
      <FormItem :label="$t('uc.finance.trade.start_end')">
        <DatePicker type="daterange" v-model="formItem.date" style="width:180px;"></DatePicker>
      </FormItem>
      <FormItem :label="$t('swap.name')">
        <Select v-model="formItem.contractCoinId" style="width:100px;" :placeholder="$t('common.pleaseselect')">
          <Option v-for="(item,index) in symbol " :value="item.id" :key="index">{{item.name}}</Option>
        </Select>
      </FormItem>
      <FormItem :label="$t('uc.finance.trade.type')">
        <Select v-model="formItem.way" style="width:100px;" :placeholder="$t('common.pleaseselect')">
          <Option value="openLong">{{$t('swap.openLong')}}</Option>
          <Option value="openShort">{{$t('swap.openShort')}}</Option>
          <Option value="closeLong">{{$t('swap.closeLong')}}</Option>
          <Option value="closeShort">{{$t('swap.closeShort')}}</Option>
        </Select>
      </FormItem>
      <!-- <FormItem :label="$t('uc.finance.trade.direction')">
        <Select v-model="formItem.direction" style="width:70px;" :placeholder="$t('common.pleaseselect')">
          <Option value="1">{{$t('uc.finance.trade.buy')}}</Option>
          <Option value="2">{{$t('uc.finance.trade.sell')}}</Option>
        </Select>
      </FormItem> -->
      <FormItem>
        <Button type="warning" @click="handleSubmit">{{$t('uc.finance.trade.search')}}</Button>
        <Button style="margin-left: 8px " @click="handleClear " class="clear_btn">{{$t('uc.finance.trade.clear')}}</Button>
      </FormItem>
    </Form>
    <div class="table">
      <Tabs :active-key="formItem.type" @on-click="setSelectedType" :animated="false">
        <Tab-pane :label="$t('swap.limited_price')" :key="2"></Tab-pane>
        <Tab-pane :label="$t('swap.trigger_price')" :key="3"></Tab-pane>
      </Tabs>
      <Table :no-data-text="$t('common.nodata')" :columns="columns" :data="orders" :loading="loading"></Table>
      <div class="page">
        <Page :total="total" :pageSize="pageSize" :current="pageNo" @on-change="loadDataPage"></Page>
      </div>
    </div>
  </div>
</template>
<script>
var moment = require("moment");
import expandRow from "@components/exchange/expand.vue";
import { swapMerge, swapSplit } from "@/utils/index";
export default {
  components: { expandRow },
  data() {
    const self = this;
    return {
      loading: false,
      pageSize: 10,
      pageNo: 1,
      total: 10,
      symbol: [],
      formItem: {
        contractCoinId: "",
        type: 2,
        entrustType: "",
        direction: "",
        date: "",
        way: ""
      },
      columns: [
          {
            title: "#",
            width: 200,
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              const way = swapMerge(row.direction, row.entrustType);
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName + (row.contractCoinType=="SECOND"?("."+row.holdTime+"SEC"):"") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  self.$t("swap."+way)
              );
            }
          },
          {
            title: self.$t("swap.time"),
            width: 200,
            key: "createTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.createTime));
            },
          },
          {
            title: self.$t("swap.priceType"),
            render(h, params) {
              return h(
                "span",
                params.row.type === "LIMIT_PRICE" ? "限价委托" : "计划委托"
              );
            },
          },
          {
            title: self.$t("swap.price"),
            key: "price",
            render(h, params) {
              return h(
                "span",
                self.toFloor(params.row.entrustPrice, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.num"),
            key: "amount",
            render(h, params) {
              return h("span", params.row.share);
            },
          },
          {
            title: self.$t("swap.strikePrice"),
            key: "strikePrice",
            render(h, params) {
              return h(
                "span",
                params.row.strikePrice > 0
                  ? self.toFloor(params.row.strikePrice, this.baseCoinScale)
                  : "--"
              );
            },
          },
          {
            title: self.$t("swap.expand.fee"),
            key: "fee",
            render(h, params) {
              return h(
                "span",
                {
                  attrs: {
                    class: 'sell',
                  },
                },
                params.row.fee==0?'--':self.toFloor(params.row.fee, this.baseCoinScale) + ' KICK'
              );
            },
          },
          {
            title: self.$t("swap.profit"),
            key: "profit",
            render(h, params) {
              const row = params.row;
              const className = row.profit >= 0 ? "buy" : "sell";
              let profit = "--"
              if ((row.entrustType == "CLOSE" && row.contractCoinType=="ALWAYS") || (row.entrustType == "OPEN" && row.contractCoinType=="SECOND")) { // 永续收益显示在平仓上  // 秒合约收益显示在开仓上
                 profit = (row.profit > 0 ? "+" : "") +
                      self.toFloor(row.profit, this.baseCoinScale);
              }
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                profit
              );
            },
          },
          {
            title: self.$t("swap.status"),
            key: "status",
            render: (h, params) => {
              const status = params.row.status;
              if (status == "ENTRUST_SUCCESS") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#f0a70a",
                    },
                  },
                  self.$t("swap._success")
                );
              } else if (status == "ENTRUST_FAILURE") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#f15057",
                    },
                  },
                  self.$t("swap.failure")
                );
              } else if (status == "ENTRUST_CANCEL") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#7c7f82",
                    },
                  },
                  self.$t("swap.canceled")
                );
              } else {
                return h("span", {}, "--");
              }
            },
          },
        ],
        columns2: [
          {
            title: "#",
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              const way = swapMerge(row.direction, row.entrustType);
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName + (row.contractCoinType=="SECOND"?("."+row.holdTime+"SEC"):"") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  self.$t("swap." + way)
              );
            }
          },
          {
            title: self.$t("swap.price"),
            key: "price",
            render(h, params) {
              return h(
                "span",
                self.toFloor(params.row.entrustPrice, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.num"),
            key: "amount",
            render(h, params) {
              return h("span", params.row.share);
            },
          },
          {
            title: self.$t("swap.triggerPrice"),
            key: "triggerPrice",
            render(h, params) {
              return h(
                "span",
                params.row.triggerPrice > 0
                  ? self.toFloor(params.row.triggerPrice, this.baseCoinScale)
                  : "--"
              );
            },
          },
          {
            title: self.$t("swap.triggerTime"),
            key: "triggeringTime",
            render: (h, params) => {
              return h("span", {}, params.row.triggeringTime?this.dateFormat(params.row.triggeringTime):'--');
            },
          },
          {
            title: self.$t("swap.strikePrice"),
            key: "strikePrice",
            render(h, params) {
              return h(
                "span",
                params.row.strikePrice > 0
                  ? self.toFloor(params.row.strikePrice, this.baseCoinScale)
                  : "--"
              );
            },
          },
          {
            title: self.$t("swap.expand.fee"),
            key: "fee",
            render(h, params) {
              return h(
                "span",
                {
                  attrs: {
                    class: 'sell',
                  },
                },
                params.row.fee==0?'--':self.toFloor(params.row.fee, this.baseCoinScale) + ' KICK'
              );
            },
          },
          {
            title: self.$t("swap.profit"),
            key: "profit",
            render(h, params) {
              const row = params.row;
              const className = row.profit >= 0 ? "buy" : "sell";
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.entrustType == "OPEN"
                  ? "--"
                  : (row.profit > 0 ? "+" : "") +
                      self.toFloor(row.profit, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.time"),
            key: "createTime",
            render: (h, params) => {
              return h("span", {}, this.dateFormat(params.row.createTime));
            },
          },
          {
            title: self.$t("swap.status"),
            key: "status",
            render: (h, params) => {
              const status = params.row.status;
              if (status == "ENTRUST_SUCCESS") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#f0a70a",
                    },
                  },
                  self.$t("swap._success")
                );
              } else if (status == "ENTRUST_FAILURE") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#f15057",
                    },
                  },
                  self.$t("swap.failure")
                );
              } else if (status == "ENTRUST_CANCEL") {
                return h(
                  "span",
                  {
                    style: {
                      color: "#7c7f82",
                    },
                  },
                  self.$t("swap.canceled")
                );
              } else {
                return h("span", {}, "--");
              }
            },
          }
        ],
      orders: []
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
    },
    "formItem.way": function (val) {
      let re = swapSplit(val);
      this.formItem.direction = re.direction;
      this.formItem.entrustType = re.entrustType;
    }
  },
  created() {
    this.refresh();
    this.getSymbol();
  },
  methods: {
    setSelectedType(e) {
      this.handleClear();
      if (e==0) { //第一个
        this.formItem.type = 2; //限价
      } else {
        this.formItem.type = 3; //计划
      }
      this.pageNo = 1;
      this.refresh();
    },
    dateFormat: function(tick) {
      return moment(tick).format("YYYY-MM-DD HH:mm:ss");
    },
    loadDataPage(data) {
      this.pageNo = data;
      this.refresh();
    },
    handleSubmit() {
      this.pageNo = 1;
      this.refresh();
    },
    handleClear() {
      this.formItem = {
        contractCoinId: "",
        direction: "",
        date: "",
        entrustType: "",
        type: this.formItem.type
      };
    },
    refresh() {
      //查询历史委托
      this.loading = true;
      const { contractCoinId, type, direction, date: rangeDate , entrustType} = this.formItem,
        startTime = new Date(rangeDate[0]).getTime() || "",
        endTime = new Date(rangeDate[1]).getTime() || "";
      let params = {};
      if (contractCoinId) params.contractCoinId = contractCoinId;
      if (direction) params.direction = direction;
      if (type) params.type = type;
      if (startTime) params.startTime = startTime;
      if (endTime) params.endTime = endTime;
      if (entrustType) params.entrustType = entrustType;
      params.pageNo = this.pageNo-1;
      params.pageSize = this.pageSize;
      var that = this;
      this.orders = [];
      this.$http
        .post(this.host + this.api.swap.history, params)
        .then(response => {
          var resp = response.body;
          if (resp.content.length > 0) {
            this.total = resp.totalElements;
            this.orders = resp.content;
          }
          this.loading = false;
        });
    },
    getSymbol() {
      this.$http.post(this.host + this.api.swap.thumb, {}).then(response => {
        var resp = response.body;
        if (resp.length > 0) {
          this.symbol = resp;
        }
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
      this.columns[9].title = this.$t("exchange.status");
    }
  }
};
</script>

