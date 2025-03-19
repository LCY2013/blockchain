<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear">
        <div class="shaow">
          <div class="money_table">
            <div style="width: 100%;height: 50px;">
              <div style="float:left;letter-spacing:1px;padding-top: 5px;">
                <span style="font-size:12px;color:#828ea1;">{{$t('uc.finance.money.totalassets')}}</span>
                <span style="font-size: 18px;color:#D8E1EB;">${{totalUSDT}}</span>
                <span style="font-size:10px;color:#828ea1;margin-left: 5px;">≈ ¥{{totalCny}}</span>
              </div>
              <Input
                style="float:right;"
                class="search"
                search
                :placeholder="$t('common.searchplaceholder')"
                @on-change="seachInputChange"
                v-model="searchKey"
              />
            </div>
            <Table
              :columns="tableColumnsMoney"
              :data="tableMoneyShow"
              :loading="loading"
              :disabled-hover="true"
            ></Table>
          </div>
        </div>
      </div>
    </div>
    <Modal v-model="transferShow" :title="$t('uc.finance.money.transfer')">
      <div class="assets-transfer">
        <div draggable="false" class="assets-transfer-inner full-screen-alert-inner secondary-font">
          <div class="transfer-info">
            <transition>
              <div class="transfer-info-left" v-if="direction">
                <div class="transfer-info-left-top">
                  <i class="transfer-info-blue"></i>
                  <p class="from-to zh-CN">{{$t('uc.finance.money.from')}}</p>
                  <p class="assets-name">{{$t('uc.finance.money.exchangeAccount')}}</p>
                </div>
                <div class="transfer-info-left-bottom">
                  <i class="transfer-info-yellow"></i>
                  <p class="from-to zh-CN">{{$t('uc.finance.money.to')}}</p>
                  <p class="assets-name">{{$t('uc.finance.money.swapAccount')}}</p>
                </div>
              </div>
              <div class="transfer-info-left" v-else>
                <div class="transfer-info-left-top">
                  <i class="transfer-info-yellow"></i>
                  <p class="from-to zh-CN">{{$t('uc.finance.money.from')}}</p>
                  <p class="assets-name">{{$t('uc.finance.money.swapAccount')}}</p>
                </div>
                <div class="transfer-info-left-bottom">
                  <i class="transfer-info-blue"></i>
                  <p class="from-to zh-CN">{{$t('uc.finance.money.to')}}</p>
                  <p class="assets-name">{{$t('uc.finance.money.exchangeAccount')}}</p>
                </div>
              </div>
            </transition>
            <div class="transfer-info-right">
              <img
                @click="changeDirection"
                src="data:image/png;base64,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"
                class="transfer-icon"
                style="transform: rotate(2160deg);"
              />
            </div>
          </div>
          <div class="input-box">
            <div class="input-box-label">{{$t('uc.finance.money.num')}}</div>
            <div class="tooltip-white tooltip-hover false" style="display: block;">
              <InputNumber
              type="text"
              class="transfer-input"
              v-model="transferAmount"
              :placeholder="$t('uc.finance.money.transfertip')"
            ></InputNumber>
              <div class="tooltip-inner-wrap">
                <div class="tooltip-inner tooltip-inner-center" style="bottom: 40px;">
                  <em class="tooltip-text">
                    <p style="text-align: center;"></p>
                  </em>
                  <i class="tooltip-arrow"></i>
                </div>
              </div>
            </div>
            <div class="input-box-currency">{{transfer.coin.name}}</div>
            <Button @click="setAmountAll" class="all transparent-button blue-font">{{$t('uc.finance.money.all')}}</Button>
          </div>
          <div style="margin-bottom: 15px;">
            <span class="max-transfer_number" style="float: left;color:#ccc;">
              ({{$t('uc.finance.money.available')}} {{transfer.coin.name}} {{direction?transfer.mainBalance:transfer.balance}})
            </span>
          </div>
        </div>
      </div>
      <div slot="footer">
        <Button type="default" size="large" @click="transferShow=false">{{$t("common.close")}}</Button>
        <Button type="primary" size="large" @click="doTransfer()">{{$t("common.ok")}}</Button>
      </div>
    </Modal>
  </div>
</template>
<script>
export default {
  components: {},
  data() {
    return {
      direction: true, //true 转入 false 转出
      modal: false,
      loginmsg: this.$t("common.logintip"),
      loading: true,
      ordKeyword: "",
      tableMoney: [],
      tableMoneyShow: [],
      searchKey: "",
      // 划转
      transferShow: false,
      transferCoins: [
        {
          value: 'BTC',
          label: 'BTC'
        }
      ],
      transferAmount: null,
      transfer: {
        coin: ''
      }
    };
  },
  methods: {
    changeDirection() {
      this.direction = !this.direction;
    },
    setAmountAll() {
      if (this.direction) { // 转入
        this.transferAmount = this.transfer.mainBalance;
      } else { // 转出
        this.transferAmount = this.transfer.balance;
      }
    },
    doTransfer() {
      if (this.direction&&(this.transferAmount>this.transfer.mainBalance)) {
        this.$Notice.error({
          title: this.$t("common.tip"),
          desc: this.$t("uc.finance.money.transferout")
        });
        return;
      }

      if (!this.direction&&(this.transferAmount>this.transfer.balance)) {
        this.$Notice.error({
          title: this.$t("common.tip"),
          desc: this.$t("uc.finance.money.transferout")
        });
        return;
      }

      this.$http
        .post(this.host + this.api.uc.contractWalletTransfer, {
          unit: this.transfer.coin.unit, 
          amount: this.transferAmount, 
          direction: this.direction?1:2
        })
        .then((response) => {
          var resp = response.body;
          if (resp.code == 0) {
            this.$Notice.success({
              title: this.$t("common.tip"),
              desc: this.$t("uc.finance.money.transfersuccess"),
            });
            this.transferShow = false;
            this.transfer = {coin: ''};
            this.getMoney();
          } else {
            this.$Notice.error({
              title: this.$t("common.tip"),
              desc: resp.message,
            });
          }
        });
    },
    seachInputChange() {
      this.tableMoneyShow = this.tableMoney.filter(
        (item) => item["coinType"].indexOf(this.searchKey) == 0
      );
    },
    getMoney() {
      //获取
      this.$http
        .post(this.host + this.api.uc.contractWallet)
        .then((response) => {
          var resp = response.body;
          if (resp.code == 0) {
            this.tableMoney = resp.data;
            for (let i = 0; i < this.tableMoney.length; i++) {
              this.tableMoney[i]["coinType"] = this.tableMoney[i].coin.unit;
            }
            this.loading = false;
            this.tableMoneyShow = this.tableMoney;
          } else {
            this.$Message.error(this.loginmsg);
          }
        });
    },
  },
  created() {
    this.getMoney();
    // this.transfer.coin = 'BTC';
  },
  computed: {
    totalUSDT() {
      let usdtTotal = 0;
      for (let i = 0; i < this.tableMoney.length; i++) {
        usdtTotal +=
          (this.tableMoney[i].balance + this.tableMoney[i].frozenBalance) *
          this.tableMoney[i].coin.usdRate;
      }
      return usdtTotal.toFixed(2);
    },
    totalCny() {
      let cnyTotal = 0;
      for (let i = 0; i < this.tableMoney.length; i++) {
        cnyTotal +=
          (this.tableMoney[i].balance + this.tableMoney[i].frozenBalance) *
          this.tableMoney[i].coin.cnyRate;
      }
      return cnyTotal.toFixed(2);
    },
    tableColumnsMoney() {
      let self = this;
      let columns = [];
      columns.push({
        title: this.$t("uc.finance.money.cointype"),
        key: "coinType",
        width: 100,
        align: "center",
      });
      columns.push({
        title: this.$t("uc.finance.money.balance"),
        key: "balance",
        align: "center",
        render(h, params) {
          return h(
            "span",
            {
              attrs: {
                title: params.row.balance,
              },
            },
            self.toFloor(params.row.balance || "0")
          );
        },
      });
      columns.push({
        title: this.$t("uc.finance.money.frozen"),
        key: "frozenBalance",
        align: "center",
        render(h, params) {
          return h(
            "span",
            {
              attrs: {
                title: params.row.frozenBalance,
              },
            },
            self.toFloor(params.row.frozenBalance || "0")
          );
        },
      });
      columns.push({
        title: this.$t("uc.finance.money.needreleased"),
        align: "center",
        render(h, params) {
          return h(
            "span",
            {
              attrs: {
                title: params.row.toReleased,
              },
            },
            self.toFloor(params.row.toReleased || "0")
          );
        },
      });
      columns.push({
        title: this.$t("uc.finance.money.operate"),
        key: "price1",
        align: "center",
        render: function (h, params) {
          var actions = [];
          actions.push(
            h(
              "Button",
              {
                // 划转;
                props: {
                  type: "info",
                  size: "small",
                },
                on: {
                  click: function () {
                    self.transfer = params.row;
                    self.transferShow = true;
                  },
                },
                style: {
                  marginRight: "0px",
                },
              },
              self.$t("uc.finance.money.transfer")
            )
          );
          return h("p", actions);
        },
      });
      return columns;
    },
  },
};
</script>
<style lang="scss">
@import url(../../assets/css/transfer.css);
// .v-enter,
// .v-leave-to{
//     opacity: 0;
//     transform: translateX(100px);
// }
// .v-enter-active,
// .v-leave-active{
//     transition: all 0.4s ease;
// }
@media screen and (max-width:1200px) {
.assets-transfer .assets-transfer-inner{
  width: 100%;
}
.assets-transfer .assets-transfer-inner .transfer-info:before{
  width: 70%;
}
}
.assets-transfer .assets-transfer-inner .input-box .all{
  position: absolute;
    padding-left: 20px;
    padding-right: 14px;
    right: 2px;
    top: 5px;
    height: 30px;
    line-height: 20px;
    font-size: 12px;
    border-left: 1px solid #CFD3E9;
    border-color: #CFD3E9;
    background-color: #192330;
    border-radius: 0;
}
.ivu-btn.ivu-btn-default{
  color:#f0ac19;
}
.ivu-modal{
  top:25%;
}
.nav-right {
  .rightarea.bill_box {
    .shaow {
      padding: 5px;
    }
    .money_table {
      .search {
        width: 200px;
        margin-bottom: 10px;
      }
      .ivu-table-wrapper {
        .ivu-table-header {
          background: #27313e;
          th {
            color: #fff;
          }
        }
        .ivu-table-body {
          td {
            color: #fff;
            .ivu-table-cell {
              padding: 10px 10px;
              p .ivu-btn {
                background: transparent;
                height: 25px;
                padding: 0 0px;
                border-radius: 0;
                span {
                  display: inline-block;
                  line-height: 20px;
                  font-size: 12px;
                  padding: 0 15px;
                  letter-spacing: 1px;
                }
              }
              p .ivu-btn.ivu-btn-info {
                border: 1px solid #f0ac19;
                span {
                  color: #f0ac19;
                }
              }
              p .ivu-btn.ivu-btn-error {
                border: 1px solid #f15057;
                span {
                  color: #f15057;
                }
              }
              p .ivu-btn.ivu-btn-primary {
                border: 1px solid #00b275;
                border: 1px solid #00b275;
                span {
                  color: #00b275;
                }
              }
              p .ivu-btn.ivu-btn-default {
                border: 1px solid #2c384f;
                background: #222c3e;
                span {
                  color: #54637a;
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>

<style scoped lang="scss">
.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 0 0 15px;
  .rightarea.bill_box {
    padding-left: 15px;
    width: 100%;
    height: auto;
    overflow: hidden;
  }
}

.demo-spin-icon-load {
  animation: ani-demo-spin 1s linear infinite;
}
@media screen and (max-width: 768px) {
  .search {
    display: none;
  }
}
</style>
