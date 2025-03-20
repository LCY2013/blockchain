<template>
  <div class="container swap" :class="skin">
    <div class="main">
      <div
        style="
          display: flex;
          flex: 100%;
          flex-wrap: wrap;
          justify-content: space-between;
        "
      >
        <div class="center">
          <div class="symbol">
            <div class="item" style="margin-left: 10px">
              <span class="coin">{{ currentCoin.name }}</span>
              <Poptip
                trigger="hover"
                :title="coinInfo.name"
                content="content"
                placement="bottom-start"
                word-wrap
                width="300"
              >
                <Icon
                  type="md-information-circle"
                  style="color: #546886; margin-left: 5px"
                />
                <div class="api" slot="content">
                  <div class="coin-info">{{ coinInfo.information }}</div>
                  <p style="text-align: right; margin-top: -10px">
                    <a :href="coinInfo.infolink" target="_blank">{{
                      $t("swap.moredetail")
                    }}</a>
                  </p>
                </div>
              </Poptip>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.NewPrice") }}</span>
              <span
                class="num"
                :class="{
                  buy: currentCoin.change > 0,
                  sell: currentCoin.change < 0,
                }"
                >{{ currentCoin.close | toFixed(4) }}</span
              >
              <span class="price-cny"
                >≈ ￥{{ (currentCoin.usdRate * CNYRate) | toFixed(2) }}</span
              >
            </div>
            <div class="item">
              <span class="text">{{ $t("service.Change") }}</span>
              <span
                class="num"
                :class="{
                  buy: currentCoin.change > 0,
                  sell: currentCoin.change < 0,
                }"
                >{{ currentCoin.rose }}</span
              >
            </div>
            <div class="item">
              <span class="text">{{ $t("service.high") }}</span>
              <span class="num">{{ currentCoin.high | toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.low") }}</span>
              <span class="num">{{ currentCoin.low | toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t("service.ExchangeNum") }}</span>
              <span class="num"
                >{{ currentCoin.volume | toFixed(4) }}
                {{ currentCoin.coin }}</span
              >
            </div>
            <div class="item">
            </div>
          </div>
          <div class="imgtable">
            <div class="handler">
              <span
                @click="changeImgTable('k')"
                :class="{ active: currentImgTable === 'k' }"
                >{{ $t("swap.kline") }}</span
              >
              <span
                @click="changeImgTable('s')"
                :class="{ active: currentImgTable === 's' }"
                >{{ $t("swap.depth") }}</span
              >
            </div>
            <div
              id="swap_kline_container"
              :class="{ hidden: currentImgTable === 's' }"
            ></div>
            <DepthGraph
              :class="{ hidden: currentImgTable === 'k' }"
              ref="depthGraph"
            ></DepthGraph>
          </div>
        </div>
        <!-- <div style="width: 100%; margin-top: 5px; flex: 0 0 100%">
          <div
            class="order"
            style="background: #999; margin-right: 5px; min-height: 352px"
          >
            <div class="order-handler">
              <span
                @click="changeOrder('currentPositions')"
                :class="{ active: selectedOrder === 'currentPositions' }"
                >{{ $t("swap.currentposition") }}</span
              >
              <span
                @click="changeOrder('current')"
                :class="{ active: selectedOrder === 'current' }"
                >{{ $t("swap.curdelegation") }}</span
              >
              <span
                @click="changeOrder('history')"
                :class="{ active: selectedOrder === 'history' }"
                >{{ $t("swap.hisdelegation") }}</span
              >
              <router-link
                v-show="selectedOrder === 'current'"
                class="linkmore"
                to="/uc/contract/entrust/current"
                >{{ $t("common.more") }}</router-link
              >
              <router-link
                v-show="selectedOrder === 'history'"
                class="linkmore"
                to="/uc/contract/entrust/history"
                >{{ $t("common.more") }}</router-link
              >
            </div>
            <div class="table">
              <Table
                class="position_table"
                height="320"
                v-if="selectedOrder === 'currentPositions'"
                :columns="currentPosition.columns"
                :data="currentPosition.rows"
                :no-data-text="$t('common.nodata')"
              ></Table>
              <div v-if="selectedOrder === 'current'">
                <div class="order-type-btn-wrap" style="height: 40px">
                  <div
                    class="order-type-btn-wrap-left"
                    style="padding-top: 5px"
                  >
                    <button
                      @click="setSelectedType(2)"
                      :class="{ 'btn-selected': selectedType == 2 }"
                    >
                      {{ $t("swap.limited_price") }}
                    </button>
                    <button
                      @click="setSelectedType(3)"
                      :class="{ 'btn-selected': selectedType == 3 }"
                    >
                      {{ $t("swap.trigger_price") }}
                    </button>
                  </div>
                </div>
                <Table
                  v-show="selectedType == 2"
                  height="280"
                  :columns="currentOrder.columns"
                  :data="currentOrder.rows"
                  :no-data-text="$t('common.nodata')"
                ></Table>
                <Table
                  v-show="selectedType == 3"
                  height="280"
                  :columns="currentOrder.columns2"
                  :data="currentOrder.rows"
                  :no-data-text="$t('common.nodata')"
                ></Table>
              </div>
              <div v-if="selectedOrder === 'history'">
                <div class="order-type-btn-wrap" style="height: 40px">
                  <div
                    class="order-type-btn-wrap-left"
                    style="padding-top: 5px"
                  >
                    <button
                      @click="setSelectedType(2)"
                      :class="{ 'btn-selected': selectedType == 2 }"
                    >
                      {{ $t("swap.limited_price") }}
                    </button>
                    <button
                      @click="setSelectedType(3)"
                      :class="{ 'btn-selected': selectedType == 3 }"
                    >
                      {{ $t("swap.trigger_price") }}
                    </button>
                  </div>
                </div>
                <Table
                  v-show="selectedType == 2"
                  height="280"
                  :columns="historyOrder.columns"
                  :no-data-text="$t('common.nodata')"
                  :data="historyOrder.rows"
                ></Table>
                <Table
                  v-show="selectedType == 3"
                  height="280"
                  :columns="historyOrder.columns2"
                  :no-data-text="$t('common.nodata')"
                  :data="historyOrder.rows"
                ></Table>
              </div>
            </div>
          </div>
        </div> -->
      </div>
      <!-- 盘口 -->
      <!-- <div class="left plate-wrap" style="position: relative; flex: 0 0 18%"> -->
        <!-- <div class="handlers">
          <span
            @click="changePlate('all')"
            class="handler handler-all"
            :class="{ active: selectedPlate == 'all' }"
          ></span>
          <span
            @click="changePlate('buy')"
            class="handler handler-green"
            :class="{ active: selectedPlate == 'buy' }"
          ></span>
          <span
            @click="changePlate('sell')"
            class="handler handler-red"
            :class="{ active: selectedPlate == 'sell' }"
          ></span>
        </div> -->
        <!-- 盘口：卖 -->
        <!-- <Table
          v-show="selectedPlate != 'buy'"
          @on-current-change="buyPlate"
          highlight-row
          ref="currentRowTable"
          class="sell_table"
          :columns="plate.columns"
          :data="plate.askRows"
          :no-data-text="$t('common.nodata')"
        ></Table> -->
        <!-- <div class="plate-nowprice">
          <span
            class="price"
            :class="{
              buy: currentCoin.change > 0,
              sell: currentCoin.change < 0,
            }"
            >{{ currentCoin.price | toFixed(baseCoinScale) }}</span
          >
          <span v-if="currentCoin.change > 0" class="buy">↑</span>
          <span v-else-if="currentCoin.change < 0" class="sell">↓</span>
          <span class="price-cny"
            >≈ {{ (currentCoin.usdRate * CNYRate) | toFixed(2) }} CNY</span
          >
        </div> -->
        <!-- 盘口：买 -->
        <!-- <Table
          v-show="selectedPlate != 'sell'"
          @on-current-change="sellPlate"
          highlight-row
          class="buy_table"
          :class="{ hidden: selectedPlate === 'all' }"
          :columns="plate.columns"
          :data="plate.bidRows"
          :no-data-text="$t('common.nodata')"
        ></Table> -->
        <!-- 开仓/平仓 -->
        <!-- <div
          class="order"
          style="margin-top: 5px; background-color: #192330; min-height: 352px"
        > -->
          <!-- <div
            v-if="currentCoin.type === 'ALWAYS'"
            class="order-handler"
            style="border-bottom: 1px solid rgb(39, 49, 62)"
          >
            <span
              :style="{
                width: currentCoin.type === 'ALWAYS' ? '50%' : '100%',
                textAlign: 'center',
              }"
              @click="entrustChange(1)"
              :class="{ active: form.entrustType === 1 }"
              >{{ $t("swap.open") }}</span
            >
            <span
              v-if="currentCoin.type === 'ALWAYS'"
              style="width: 50%; text-align: center"
              @click="entrustChange(2)"
              :class="{ active: form.entrustType === 2 }"
              >{{ $t("swap.close") }}</span
            >
          </div> -->
          <!-- <div class="table open-close">
            <div class="open" style="text-align: center; margin-top: 10px"> -->
              <!-- <RadioGroup v-model="form.type" type="button" size="default">
                <Radio label="2">{{ $t("swap.limited_price") }}</Radio>
                <Radio label="3">{{ $t("swap.trigger_price") }}</Radio>
              </RadioGroup> -->
              <!-- <div class="account-item">
                <div style="width:35%;" >{{$t("swap.accountmode")}}</div>     
                <div class="mode" style="width: 65%;">
                  <Button @click="showMarginModeModal()" v-if="form.patterns === '1'" size="small">{{$t("swap.marginMode2")}}</Button>
                  <Button @click="showMarginModeModal()" v-if="form.patterns === '2'" size="small">{{$t("swap.marginMode1")}}</Button>
                </div>
              </div>-->
              <!-- <div
                class="account-item"
                style="
                  width: 90%;
                  margin-top: 10px;
                  height: 32px;
                  padding-bottom: 5px;
                "
              >
                <div style="width: 35%; text-align: left">
                  {{ $t("swap.accountmargin") }}
                </div>
                <div
                  class="margin"
                  style="
                    width: 65%;
                    justify-content: space-between;
                    text-align: right;
                  "
                >
                  <Button
                    @click="showAdjustLeverage(1)"
                    size="small"
                    style="
                      flex: 0 0 47%;
                      border: none;
                      background: rgb(240 172 25);
                      color: #000;
                      padding: 0px 10px;
                    "
                    >{{ form.leverage }}X</Button
                  >
                </div>
              </div> -->
              <!-- <div
                class="account-item"
                style="
                  width: 90%;
                  margin-top: 10px;
                  height: 32px;
                  padding-bottom: 5px;
                "
                v-if="currentCoin.type === 'SECOND'"
              >
                <div style="width: 35%; text-align: left">
                  {{ $t("swap.time") }}
                </div>
                <div
                  class="margin"
                  style="
                    width: 65%;
                    justify-content: space-between;
                    text-align: right;
                  "
                >
                  <Button
                    @click="showAdjustTime(1)"
                    size="small"
                    style="
                      flex: 0 0 47%;
                      border: none;
                      background: rgb(240 172 25);
                      color: #000;
                      padding: 0px 10px;
                    "
                    >{{ form.time }}S</Button
                  >
                </div>
              </div> -->
              <!-- <Form style="width: 90%; margin: 0 auto; margin-top: 18px">
                <FormItem style="margin-bottom: 18px" v-if="form.type == 3">
                  <label
                    class="before"
                    style="
                      float: left;
                      color: #b0b8db;
                      min-width: 58px;
                      text-align: left;
                    "
                    >{{ $t("swap.triggerPrice") }}</label
                  >
                  <InputNumber
                    style="float: left; width: calc(100% - 60px)"
                    @on-keyup="keyEvent"
                    v-model="form.triggerPrice"
                  >
                  </InputNumber>
                </FormItem>
                <FormItem style="margin-bottom: 18px">
                  <label
                    class="before"
                    style="
                      float: left;
                      color: #b0b8db;
                      min-width: 58px;
                      text-align: left;
                    "
                    >{{ $t("swap.price") }}</label
                  >
                  <Button
                    v-if="form.market"
                    @click="form.market = !form.market"
                    type="primary"
                    style="float: right; margin-top: 3px; padding: 3px 6px"
                    shape="circle"
                    class="operate_btn"
                    >{{ $t("swap.price") }}</Button
                  >
                  <Button
                    v-else
                    @click="form.market = !form.market"
                    type="primary"
                    style="float: right; margin-top: 3px; padding: 3px 6px"
                    shape="circle"
                    class="operate_btn"
                    >{{ $t("swap.marketPrice") }}</Button
                  >
                  <InputNumber
                    v-if="!form.market"
                    style="float: left; min-width: 90px"
                    @on-keyup="keyEvent"
                    v-model="form.limitPrice"
                  >
                  </InputNumber>
                  <Input
                    style="float: left; width: auto"
                    v-else
                    :value="$t('swap.marketPrice')"
                    disabled
                  />
                </FormItem>
                <FormItem style="margin-bottom: 18px">
                  <label
                    class="before"
                    style="
                      float: left;
                      color: #b0b8db;
                      min-width: 58px;
                      text-align: left;
                    "
                    >{{ $t("swap.num") }}</label
                  >
                  <InputNumber
                    style="float: left; width: calc(100% - 60px)"
                    @on-keyup="keyEvent"
                    v-model="form.limitAmount"
                    :min="1"
                  >
                  </InputNumber>
                </FormItem>
              </Form> -->
              <!-- <template v-if="isLogin">
                <template v-if="form.entrustType == 1">
                  <div class="operate" style="width: 70%; margin: 0 auto">
                    <div
                      class="operate_details operate_left"
                      style="float: left"
                    >
                      <div class="operate_details1">
                        <span class="green">{{ $t("swap.canup") }}:</span>
                        <span class="num">{{ maxOpen }}</span>
                      </div>
                      <Button
                        @click="openOrder(1)"
                        type="primary"
                        shape="circle"
                        class="open-up operate_btn"
                        >{{ $t("swap.openup") }}</Button
                      >
                    </div>
                    <div
                      class="operate_details operate_right"
                      style="float: right"
                    >
                      <div class="operate_details2">
                        <span class="red">{{ $t("swap.candown") }}:</span>
                        <span class="num">{{ maxOpen }}</span>
                      </div>
                      <Button
                        @click="openOrder(2)"
                        type="primary"
                        shape="circle"
                        class="open-down operate_btn"
                        >{{ $t("swap.opendown") }}</Button
                      >
                    </div>
                  </div>
                </template>
                <template v-else>
                  <div class="operate" style="width: 70%; margin: 0 auto">
                    <div
                      class="operate_details operate_left"
                      style="float: left"
                    >
                      <div class="operate_details1">
                        <span class="red">{{ $t("swap.canclosedown") }}:</span>
                        <span class="num">{{ maxCloseSellAmount }}</span>
                      </div>
                      <Button
                        @click="closeOrder(2)"
                        type="primary"
                        shape="circle"
                        class="open-down operate_btn"
                        >{{ $t("swap.closedown") }}</Button
                      >
                    </div>
                    <div
                      class="operate_details operate_right"
                      style="float: right"
                    >
                      <div class="operate_details2">
                        <span class="green">{{ $t("swap.cancloseup") }}:</span>
                        <span class="num">{{ maxCloseBuyAmount }}</span>
                      </div>
                      <Button
                        @click="closeOrder(1)"
                        type="primary"
                        shape="circle"
                        class="open-up operate_btn"
                        >{{ $t("swap.closeup") }}</Button
                      >
                    </div>
                  </div>
                </template>
              </template> -->
              <!-- <div
                class="operate-login"
                v-if="!isLogin"
                style="width: 94%; margin-left: 3%"
              >
                <span
                  style="
                    display: inline-block;
                    width: 100%;
                    text-align: center;
                    border: 1px solid #232d3a;
                    padding: 5px 0;
                    border-radius: 5px;
                  "
                >
                  {{ $t("common.please") }}
                  <router-link to="/login">
                    <span style="color: #f0a70a">{{
                      $t("common.login")
                    }}</span> </router-link
                  >/
                  <router-link to="/register">
                    <span style="color: #00dcff">{{
                      $t("common.register")
                    }}</span>
                  </router-link>
                </span>
              </div> -->
            <!-- </div>
          </div> -->
        <!-- </div> -->
      <!-- </div> -->
      <!-- 成交记录 -->
      <!-- <div class="left plate-wrap" style="position: relative; flex: 0 0 13%"> -->
        <!-- <div
          style="
            background-color: #192330;
            height: 40px;
            line-height: 40px;
            padding-left: 5px;
            color: #61688a;
            font-size: 13px;
          "
        >
          <span>{{ $t("swap.latestdeal") }}</span>
        </div> -->
        <!-- <div class="trade-wrap">
          <Table
            height="472"
            :columns="trade.columns"
            :data="trade.rows"
            :no-data-text="$t('common.nodata')"
          ></Table>
        </div> -->
        <!-- 我的合约账户 -->
        <!-- <div
          class="order"
          style="
            margin-top: 5px;
            min-height: 352px;
            background-color: #192330;
            color: #61688a;
          "
        >
          <div
            style="
              height: 33px;
              line-height: 33px;
              padding-left: 10px;
              border-bottom: 1px solid #27313e;
              font-size: 14px;
            "
          >
            <span style="color: #fff">{{ $t("swap.myswapaccount") }}</span>
            <router-link
              class="linkmore"
              to="/uc/contract-money"
              style="margin-right: 10px"
              >{{ $t("swap.zijinhuazhuan") }}</router-link
            >
          </div>
          <div class="table swap-my-account">
            <div class="account-item">
              <div>{{ $t("swap.accountquanyi") }}</div>
              <div style="font-size: 12px">
                <span
                  >{{
                    (wallet.base + wallet.frozen) | toFixed(baseCoinScale)
                  }}(USD)</span
                >
              </div>
              <div style="clear: both"></div>
            </div>
            <div class="account-item">
              <div>{{ $t("swap.profitandloss") }}</div>
              <div style="font-size: 12px">
                <span>{{ wallet.profit | toFixed(baseCoinScale) }}</span>
              </div>
              <div style="clear: both"></div>
            </div>
            <div class="account-item">
              <div>{{ $t("swap.principalAmount") }}</div>
              <div style="font-size: 12px">
                <span>{{ wallet.base | toFixed(baseCoinScale) }}(USD)</span>
              </div>
              <div style="clear: both"></div>
            </div>
            <div class="account-item">
              <div>{{ $t("swap.positionAmount") }}</div>
              <div style="font-size: 12px">
                <span>{{ positionNum }}</span>
              </div>
              <div style="clear: both"></div>
            </div>
            <div class="account-item">
              <div>{{ $t("swap.frozenAmount") }}</div>
              <div style="font-size: 12px">
                <span>{{ wallet.frozen | toFixed(baseCoinScale) }}(USD)</span>
              </div>
              <div style="clear: both"></div>
            </div>
          </div>
        </div> -->
      <!-- </div> -->
    </div>

    <!-- 弹出框: 变更仓位模式 -->
    <!-- <Modal
      v-model="marginModeModal"
      :title="$t('swap.modifyMarginModeTitle')"
      width="20"
      class-name="vertical-center-modal"
    >
      <div style="text-align: center; width: 100%">
        <ButtonGroup style="width: 80%; text-align: center">
          <Button
            style="width: 50%"
            v-if="form.patterns === '1'"
            type="primary"
            >{{ $t("swap.marginMode1") }}</Button
          >
          <Button style="width: 50%" v-else @click="changeMarginMode('1')">{{
            $t("swap.marginMode1")
          }}</Button>

          <Button
            style="width: 50%"
            v-if="form.patterns === '2'"
            type="primary"
            >{{ $t("swap.marginMode2") }}</Button
          >
          <Button style="width: 50%" v-else @click="changeMarginMode('2')">{{
            $t("swap.marginMode2")
          }}</Button>
        </ButtonGroup>
      </div>
      <div slot="footer">
        <Button type="default" size="large" @click="marginModeModal = false">{{
          $t("common.close")
        }}</Button>
      </div>
    </Modal> -->

    <!-- 弹出框: 调整时间 -->
    <Modal
      v-model="form.timeModal"
      :title="$t('swap.time')"
      width="20"
      class-name="vertical-center-modal"
    >
      <div class="leverage">
        <Slider
          v-model="form.timeAdjustVal2"
          @on-input="xxlever2(form.timeAdjustVal2)"
          @on-change="xxlever2(form.timeAdjustVal2)"
          :marks="times"
          show-tip="never"
          :min="1"
          :max="6"
        ></Slider>
        <!-- <Select v-model="form.timeAdjustVal">
          <Option
            v-for="(item, index) in times"
            :key="index"
            :value="item.value"
          >{{ item.label }}</Option>
        </Select> -->
      </div>
      <div slot="footer">
        <Button type="default" size="large" @click="form.timeModal = false">{{
          $t("common.close")
        }}</Button>
        <Button type="primary" size="large" @click="adjustTime()">{{
          $t("common.ok")
        }}</Button>
      </div>
    </Modal>

    <!-- 弹出框: 调整杠杆 -->
    <Modal
      v-model="form.leverageModal"
      :title="$t('swap.modifyLeverage')"
      width="20"
      class-name="vertical-center-modal"
    >
      <div class="leverage">
        <Slider
          v-model="form.leverageAdjustVal2"
          @on-input="xxlever(form.leverageAdjustVal2)"
          @on-change="xxlever(form.leverageAdjustVal2)"
          :marks="leverages"
          show-tip="never"
          :min="1"
          :max="5"
        ></Slider>
        <!-- <Select v-model="form.leverageAdjustVal">
          <Option
            v-for="(item, index) in leverages"
            :key="index"
            :value="item.value"
          >{{ item.label }}</Option>
        </Select> -->
      </div>
      <div slot="footer">
        <Button
          type="default"
          size="large"
          @click="form.leverageModal = false"
          >{{ $t("common.close") }}</Button
        >
        <Button type="primary" size="large" @click="adjustLeverage()">{{
          $t("common.ok")
        }}</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import expandRow from "@components/exchange/expand.vue";
import Datafeeds from "@js/charting_library/datafeed/swaptrade.js";
var Stomp = require("stompjs");
var SockJS = require("sockjs-client");
var moment = require("moment");
import DepthGraph from "@components/exchange/DepthGraph.vue";
import $ from "@js/jquery.min.js";

export default {
  components: { expandRow, DepthGraph },
  data() {
    let self = this;
    return {
      contractId: null,
      marginModeModal: false,
      accountMode: 0,
      sliderStep: [25, 50, 75, 100],
      sliderBuyLimitPercent: 0,
      sliderSellLimitPercent: 0,
      sliderBuyMarketPercent: 0,
      sliderSellMarketPercent: 0,
      memberRate: 0,
      // userRealVerified: false, //是否实名认证
      collecRequesting: false,
      currentCoinIsFavor: false,
      isUseBHB: false,
      skin: "night", //皮肤样式day&night
      currentImgTable: "k",
      selectedType: "2", //当前显示的委托记录 委托类型 1市价 2限价 3计划委托
      selectedOrder: "currentPositions", //当前显示的委托记录
      selectedPlate: "all", //当前显示的买卖盘
      CNYRate: null,
      datafeed: null,
      defaultPath: 1,
      basecion: "usdt",
      coinScale: 6,
      baseCoinScale: 4,
      // symbolFee: 0.001,
      takerFee: 0.001,
      makerFee: 0.001,
      dataIndex: [],
      dataIndex2: [],
      searchKey: "",
      coinInfo: {},
      currentCoin: {
        id: "",
        base: "",
        coin: "",
        symbol: "",
        perUsdt: 100000,
        close: 0,
      },
      leverages: {
        1: "1x",
        2: "2x",
        3: "5x",
        4: "10x",
        5: "20x",
      },
      times: {
        1: "10S",
        2: "30S",
        3: "60S",
        4: "120S",
        5: "360S",
        6: "3000S",
      },
      enableMarketBuy: 1, // 0:禁用  1:启用
      enableMarketSell: 1,
      exchangeable: 1, // 0:可交易   1:不可交易
      //当前市场上的交易币种，按交易对分
      coins: {
        _map: [],
        USDT: [],
        BTC: [],
        ETH: [],
        USDT2: [],
        favor: [],
        columns: [
          {
            title: this.$t("swap.contract") + "/" + this.$t("swap.vol"),
            key: "coin",
            minWidth: 40,
            sortable: false,
            className: "coin-menu-symbol swap-coin-menu-symbol",
            render: (h, params) => {
              return h(
                "div",
                {
                  style: {
                    padding: "5px 0",
                  },
                },
                [
                  h(
                    "span",
                    {
                      style: {
                        fontSize: "14px",
                      },
                    },
                    params.row.name
                  ),
                  h("br"),
                  h(
                    "span",
                    {
                      style: {
                        color: "#61688A",
                        fontSize: "14px",
                      },
                    },
                    params.row.volume.toFixed(4)
                  ),
                ]
              );
            },
          },
          {
            title: this.$t("swap.lastprice") + "/" + this.$t("swap.daychange"),
            key: "rose",
            minWidth: 40,
            className: "swap-coin-menu-lastprice",
            sortMethod: function (a, b, type) {
              let a1 = a.replace(/[^\d|.|-]/g, "") - 0;
              let b1 = b.replace(/[^\d|.|-]/g, "") - 0;
              if (type == "asc") {
                return a1 - b1;
              } else {
                return b1 - a1;
              }
            },
            render: (h, params) => {
              const row = params.row;
              const className = parseFloat(row.rose) < 0 ? "sell" : "buy";
              return h(
                "div",
                {
                  style: {
                    padding: "5px 10px",
                    textAlign: "right",
                    fontSize: "14px",
                  },
                },
                [
                  h("span", params.row.close),
                  h("br"),
                  h(
                    "span",
                    {
                      attrs: {
                        class: className,
                      },
                    },
                    row.rose
                  ),
                ]
              );
            },
          },
        ],
      },
      wallet: {
        base: 0.0,
        frozen: 0.0,
        profit: 0.0,
        kick: 0.0,
      },
      showMarket: false,
      fixHistoryHeight: 295,
      // rechargeUrl:'#/finance/recharge',
      // rechargeUSDTUrl:'#/finance/recharge?name=USDT',
      maxCloseSellAmount: "-",
      maxCloseBuyAmount: "-",
      // times: [
      //   {
      //     value: 10,
      //     label: 10,
      //   },
      //   {
      //     value: 30,
      //     label: 30,
      //   },
      //   {
      //     value: 60,
      //     label: 60,
      //   },
      //   {
      //     value: 120,
      //     label: 120,
      //   },
      //   {
      //     value: 360,
      //     label: 360,
      //   },
      //   {
      //     value: 3000,
      //     label: 3000,
      //   }
      // ],
      form: {
        type: "2", //委托类型 1市价 2限价 3计划委托
        time: 10, //时间
        patterns: "1", //1全仓 2逐仓
        entrustType: 1, // 1 开仓 2 平仓 ｜ 与 direction 配合（）
        limitPrice: 0.0, //价格
        triggerPrice: 0, //触发价
        limitAmount: 0, //张数
        marketAmount: 0.0,
        limitTurnover: 0.0,
        limitPercent: 0, //百分比
        leverageModal: false,
        leverage: 10,
        timeAdjustVal: 10,
        timeAdjustVal2: 1,
        leverageAdjustVal: 10,
        leverageAdjustVal2: 4,
        market: false,
      },
      trade: {
        rows: [],
        columns: [
          {
            title: self.$t("swap.price"),
            key: "price",
            render: (h, params) => {
              const row = params.row;
              const className = row.direction == "BUY" ? "buy" : "sell";
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                params.row.price
              );
            },
            renderHeader: (h, params) => {
              const title =
                self.$t("swap.price") + "(" + self.currentCoin.base + ")";
              return h("span", {}, title);
            },
          },
          {
            title: self.$t("swap.num"),
            key: "amount",
            render: (h, params) => {
              return h("span", {}, params.row.amount.toFixed(4));
            },
            renderHeader: (h, params) => {
              const title = self.$t("swap.num"); // + "(" + self.currentCoin.coin + ")"
              return h("span", {}, title);
            },
          },
          {
            title: self.$t("swap.time"),
            key: "time",
            render: (h, params) => {
              return h("span", {}, this.timeFormat(params.row.time));
            },
          },
        ],
      },
      //   最新价的 table 数据;
      plate: {
        maxPostion: 10,
        columns: [
          {
            title: self.$t("swap.price"),
            key: "price",
            render: (h, params) => {
              let str = "";
              let price = "";
              const className = params.row.direction.toLowerCase();
              params.row.price == 0 && (str = h("span", {}, "--"));
              params.row.price != 0 &&
                (price = params.row.price.toFixed(4)) &&
                (str = h(
                  "span",
                  {
                    attrs: {
                      class: className,
                    },
                  },
                  price
                ));
              return str;
            },
            renderHeader: (h, params) => {
              const title =
                self.$t("swap.price") + "(" + self.currentCoin.base + ")";
              return h("span", {}, title);
            },
          },
          {
            title: self.$t("swap.num"),
            key: "amount",
            render: (h, params) => {
              let str = "";
              params.row.amount == 0 && (str = h("span", {}, "--"));
              params.row.amount != 0 &&
                (str = h("span", {}, params.row.amount.toFixed(4)));
              return str;
            },
            renderHeader: (h, params) => {
              const title = self.$t("swap.num");
              // + "(" + self.currentCoin.coin + ")"
              return h("span", {}, title);
            },
          },
          {
            title: self.$t("swap.total"),
            key: "totalAmount",
            render: (h, params) => {
              if (params.row.price == 0 || params.row.totalAmount == 0) {
                return h("span", {}, "--");
              } else {
                return h("span", {}, params.row.totalAmount.toFixed(4));
              }
            },
            renderHeader: (h, params) => {
              const title =
                self.$t("swap.total") + "(" + self.currentCoin.coin + ")";
              return h("span", {}, title);
            },
          },
          {
            className: "percenttd",
            width: 1,
            render: (h, params) => {
              let width = "0",
                backgroundColor =
                  params.row.direction === "BUY" ? "#00b275" : "#f15057",
                totle =
                  params.row.direction === "BUY"
                    ? this.plate.bidTotle
                    : this.plate.askTotle;
              if (params.row.totalAmount) {
                width = (params.row.totalAmount / totle).toFixed(4) * 100 + "%";
              }
              return h(
                "div",
                {
                  style: {
                    width,
                    backgroundColor,
                  },
                  attrs: {
                    class: "percentdiv",
                  },
                },
                " "
              );
            },
          },
        ],
        askRows: [],
        bidRows: [],
      },
      currentPosition: {
        columns: [
          {
            title: "#",
            width: 180,
            height: 40,
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName +
                  (row.contractCoinType == "SECOND"
                    ? "." + row.holdTime + "SEC"
                    : "") +
                  "." +
                  row.multiple +
                  "X" +
                  "." +
                  (row.direction == "BUY"
                    ? self.$t("swap.up")
                    : self.$t("swap.down"))
              );
            },
          },
          // {
          //   title: self.$t("swap.time"),
          //   key: "time",
          //   render: (h, params) => {
          //     return h("span", {}, this.dateFormat(params.row.time));
          //   }
          // },
          {
            title: self.$t("swap.avgOpenPrice"),
            key: "openPrice",
            render: (h, params) => {
              return h("span", {}, params.row.openPrice.toFixed(2));
            },
          },
          {
            title: self.$t("swap.profitandloss"),
            key: "profit",
            render: (h, params) => {
              const row = params.row;
              const className = row.profit > 0 ? "buy" : "sell";
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.profit
              );
            },
          },
          // {
          //   title: self.$t("swap.priceType"),
          //   render(h, params) {
          //     return h(
          //       "span",
          //       params.row.type === "LIMIT_PRICE" ? "限价委托" : "计划委托"
          //     );
          //   }
          // },
          // {
          //   title: self.$t("swap.direction"),
          //   key: "direction",
          //   render: (h, params) => {
          //     const row = params.row;
          //     const className = row.direction.toLowerCase();
          //     return h(
          //       "span",
          //       {
          //         attrs: {
          //           class: className
          //         }
          //       },
          //       row.direction == "BUY"
          //         ? (row.entrustType=="OPEN"?self.$t("swap.openLong"):self.$t("swap.closeShort"))
          //         : (row.entrustType=="OPEN"?self.$t("swap.openShort"):self.$t("swap.closeLong"))
          //     );
          //   }
          // },
          {
            title: self.$t("swap.frozenAmount"),
            key: "principalAmount",
            render(h, params) {
              return h(
                "span",
                self.toFloor(params.row.principalAmount, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.amount"),
            key: "position",
          },
          {
            title: self.$t("swap.closeAble"),
            key: "availablePosition",
          },
          {
            title: self.$t("swap.liquidationPrice"),
            key: "liquidationPrice",
            render(h, params) {
              return h("span", self.toFloor(params.row.liquidationPrice, 2));
            },
          },
          {
            title: self.$t("swap.action"),
            key: "operate",
            width: 110,
            render: (h, params) => {
              const row = params.row;
              let buttons = [];
              if (row.contractCoinType == "SECOND") {
                buttons = [];
              } else {
                buttons = [
                  row.direction == "BUY"
                    ? h(
                        "Button",
                        {
                          props: {
                            size: "small",
                          },
                          attrs: {
                            class: "bg-red",
                          },
                          on: {
                            click: () => {
                              // console.log("======开始撤单")
                              this.quickClose(row);
                            },
                          },
                        },
                        self.$t("swap.quickclose")
                      )
                    : h(
                        "Button",
                        {
                          props: {
                            size: "small",
                          },
                          attrs: {
                            class: "bg-green",
                          },
                          on: {
                            click: () => {
                              // console.log("======开始撤单")
                              this.quickClose(row);
                            },
                          },
                        },
                        self.$t("swap.quickclose")
                      ),
                ];
              }
              return h("div", buttons);
            },
          },
        ],
        rows: [],
      },
      currentOrder: {
        columns: [
          {
            title: "#",
            width: 180,
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName +
                  (row.contractCoinType == "SECOND"
                    ? "." + row.holdTime + "SEC"
                    : "") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  (row.direction == "BUY"
                    ? row.entrustType == "OPEN"
                      ? self.$t("swap.openLong")
                      : self.$t("swap.closeShort")
                    : row.entrustType == "OPEN"
                    ? self.$t("swap.openShort")
                    : self.$t("swap.closeLong"))
              );
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
                params.row.marketPrice
                  ? "对手价"
                  : self.toFloor(params.row.entrustPrice, self.baseCoinScale)
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
                  : "-"
              );
            },
          },
          {
            title: self.$t("swap.frozenAmount"),
            key: "principalAmount",
            render(h, params) {
              return h(
                "span",
                self.toFloor(params.row.principalAmount, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.action"),
            key: "operate",
            width: 110,
            render: (h, params) => {
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
                      this.cancel(params.index);
                    },
                  },
                },
                self.$t("swap.undo")
              );
            },
          },
        ],
        columns2: [
          {
            title: "#",
            width: 180,
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName +
                  (row.contractCoinType == "SECOND"
                    ? "." + row.holdTime + "SEC"
                    : "") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  (row.direction == "BUY"
                    ? row.entrustType == "OPEN"
                      ? self.$t("swap.openLong")
                      : self.$t("swap.closeShort")
                    : row.entrustType == "OPEN"
                    ? self.$t("swap.openShort")
                    : self.$t("swap.closeLong"))
              );
            },
          },
          {
            title: self.$t("swap.price"),
            key: "price",
            render(h, params) {
              return h(
                "span",
                params.row.marketPrice
                  ? "对手价"
                  : self.toFloor(params.row.entrustPrice, self.baseCoinScale)
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
            title: self.$t("swap.strikePrice"),
            key: "strikePrice",
            render(h, params) {
              return h(
                "span",
                params.row.strikePrice > 0
                  ? self.toFloor(params.row.strikePrice, this.baseCoinScale)
                  : "-"
              );
            },
          },
          {
            title: self.$t("swap.frozenAmount"),
            key: "principalAmount",
            render(h, params) {
              return h(
                "span",
                self.toFloor(params.row.principalAmount, this.baseCoinScale)
              );
            },
          },
          {
            title: self.$t("swap.action"),
            key: "operate",
            width: 110,
            render: (h, params) => {
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
                      this.cancel(params.index);
                    },
                  },
                },
                self.$t("swap.undo")
              );
            },
          },
        ],
        rows: [],
      },
      historyOrder: {
        pageSize: 10,
        total: 10,
        page: 0,
        columns: [
          {
            title: "#",
            width: 180,
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName +
                  (row.contractCoinType == "SECOND"
                    ? "." + row.holdTime + "SEC"
                    : "") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  (row.direction == "BUY"
                    ? row.entrustType == "OPEN"
                      ? self.$t("swap.openLong")
                      : self.$t("swap.closeShort")
                    : row.entrustType == "OPEN"
                    ? self.$t("swap.openShort")
                    : self.$t("swap.closeLong"))
              );
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
                params.row.marketPrice
                  ? "对手价"
                  : self.toFloor(params.row.entrustPrice, self.baseCoinScale)
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
                    class: "sell",
                  },
                },
                !params.row.fee
                  ? "--"
                  : -self.toFloor(params.row.fee, this.baseCoinScale)
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
          },
        ],
        columns2: [
          {
            width: 180,
            title: "#",
            render: (h, params) => {
              const row = params.row;
              const className = row.direction.toLowerCase();
              return h(
                "span",
                {
                  attrs: {
                    class: className,
                  },
                },
                row.contractCoinName +
                  (row.contractCoinType == "SECOND"
                    ? "." + row.holdTime + "SEC"
                    : "") +
                  "." +
                  row.leverage +
                  "X" +
                  "." +
                  (row.direction == "BUY"
                    ? row.entrustType == "OPEN"
                      ? self.$t("swap.openLong")
                      : self.$t("swap.closeShort")
                    : row.entrustType == "OPEN"
                    ? self.$t("swap.openShort")
                    : self.$t("swap.closeLong"))
              );
            },
          },
          {
            title: self.$t("swap.price"),
            key: "price",
            render(h, params) {
              return h(
                "span",
                params.row.marketPrice
                  ? "对手价"
                  : self.toFloor(params.row.entrustPrice, self.baseCoinScale)
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
              return h(
                "span",
                {},
                params.row.triggeringTime
                  ? this.dateFormat(params.row.triggeringTime)
                  : "--"
              );
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
                    class: "sell",
                  },
                },
                !params.row.fee
                  ? "--"
                  : -self.toFloor(params.row.fee, this.baseCoinScale)
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
          },
        ],
        rows: [],
      },
      fullTrade: {},
    };
  },
  filters: {},
  computed: {
    positionNum: function () {
      let num = 0;
      let buyNum = 0;
      let sellNum = 0;
      this.currentPosition.rows.forEach((e) => {
        if (e.direction === "BUY") {
          //买入
          buyNum += e.availablePosition;
        } else {
          // 卖出
          sellNum += e.availablePosition;
        }
        num += e.position;
      });
      // console.log(buyNum, sellNum)
      this.maxCloseSellAmount = buyNum;
      this.maxCloseBuyAmount = sellNum;
      return num;
    },
    maxOpen: function () {
      //最高买入开多
      var balance = this.wallet.base,
        perUsdt = this.currentCoin.perUsdt,
        leverage = this.form.leverage;
      if (balance <= 0 || !perUsdt || !this.form.leverage) {
        return 0;
      }
      return parseInt((balance * leverage) / perUsdt);
    },
    rechargeUSDTUrl: function () {
      return "/uc/recharge?name=" + this.currentCoin.base;
      // return "#/finance/recharge?name=" + this.currentCoin.base;
    },
    rechargeCoinUrl: function () {
      return "/uc/recharge?name=" + this.currentCoin.coin;
      // return "#/finance/recharge?name=" + this.currentCoin.coin;
    },
    isLogin: function () {
      return this.$store.getters.isLogin;
    },
    member: function () {
      return this.$store.getters.member;
    },
    lang: function () {
      return this.$store.state.lang;
    },
    sliderBuyDisabled() {
      let account = this.wallet.base,
        min = this.toFloor(1 / Math.pow(10, this.baseCoinScale));
      return account < min;
    },
    sliderSellDisabled() {
      let account = this.wallet.base,
        min = this.toFloor(1 / Math.pow(10, this.coinScale));
      return account < min;
    },
  },
  watch: {
    lang: function () {
      this.updateLangData();
    },
    $route(to, from) {
      this.init();
    },
    "currentCoin.close": function (p) {
      let profit = 0.0;
      this.currentPosition.rows.forEach((e) => {
        console.log(e)
        if (e.direction == "BUY") {
          //多仓
          e.profit = parseFloat(
            (1 / e.openPrice - 1 / p) * e.multiple * e.principalAmount * p
          ).toFixed(this.baseCoinScale);
        } else {
          //空仓
          e.profit = parseFloat(
            (1 / p - 1 / e.openPrice) * e.multiple * e.principalAmount * p
          ).toFixed(this.baseCoinScale);
        }
        profit += parseFloat(e.profit);
      });
      this.wallet.profit = parseFloat(profit).toFixed(this.baseCoinScale);
    },
  },
  created: function () {
    this.init();
  },
  methods: {
    silderGo(silder, val) {
      this[silder] = val;
    },
    init() {
      var id = this.$route.params.pair;
      if (id == undefined) {
        this.$router.push("/swapindex/" + this.defaultPath);
        // params = this.defaultPath;
      }
      this.$Loading.start();
      this.$http
        .post(this.host + this.api.swap.symbolInfo, {
          id: id,
        })
        .then((response) => {
          var resp = response.body;
          if (!resp) {
            return;
          }
          this.basecion = resp.baseSymbol.toLowerCase();
          this.currentCoin.symbol = resp.coinSymbol + "/" + resp.baseSymbol;
          this.currentCoin.coin = resp.coinSymbol;
          this.currentCoin.id = id;
          this.currentCoin.type = resp.type;
          this.currentCoin.base = resp.baseSymbol;
          this.$store.commit("setSkin", this.skin);
          // 一些附加信息
          this.currentCoin.coinScale = resp.coinScale;
          this.currentCoin.baseCoinScale = resp.baseCoinScale;
          this.currentCoin.perUsdt = resp.shareNumber;
          this.baseCoinScale = resp.baseCoinScale;
          this.coinScale = resp.coinScale;
          this.takerFee = resp.takerFee;
          this.makerFee = resp.makerFee;
          this.enableMarketBuy = resp.enableMarketBuy;
          this.enableMarketSell = resp.enableMarketSell;
          this.exchangeable = resp.exchangeable;
          // 一些附加信息
          this.getCNYRate();
          this.getCoinInfo();
          this.getSymbol(); //包含 K线图、getFavor、startWebsock等
          this.getPlateFull();
          if (this.isLogin) {
            this.getWallet(); //账户资产信息
            this.getCurrentPosition(); //当前持仓
            this.getCurrentOrder(); //当前委托
            this.getHistoryOrder(); //历史委托
            this.getLeverage();
          }
          this.sliderBuyLimitPercent = 0;
          this.sliderSellLimitPercent = 0;
          this.sliderBuyMarketPercent = 0;
          this.form.entrustType = 1;
          this.$Loading.finish();
        });
    },
    tipFormat(val) {
      return val + "%";
    },
    changeBaseCion(str) {
      this.basecion = str;
      this.dataIndex = this.coins.USDT;
      this.dataIndex2 = this.coins.USDT2;
    },
    changePlate(str) {
      if (str != "all") {
        this.plate.maxPostion = 20;
      } else {
        this.plate.maxPostion = 10;
      }
      this.getPlate(str);
      //this.selectedPlate = str;
    },
    changeImgTable(str) {
      this.currentImgTable = str;
    },
    changeOrder(str) {
      this.selectedOrder = str;
    },
    entrustChange(val) {
      this.form.entrustType = val;
    },
    setback() {
      var obk = document.getElementsByClassName("container")[0];
      var height = 0;
      var doc = document;
      window.innerHeight && (height = window.innerHeight);
      !window.innerHeight &&
        doc.body.clientHeight &&
        (height = doc.body.clientHeight);
      !window.innerHeight &&
        !doc.body.clientHeight &&
        doc.documentElement.clientHeight &&
        (height = doc.documentElement.clientHeight);
      obk.style.minHeight = height - 100 + "px";
    },
    updateLangData() {
      this.coins.columns[0].title = this.$t("swap.coin");
      this.coins.columns[1].title = this.$t("swap.lastprice");
      this.coins.columns[2].title = this.$t("swap.daychange");
      // this.coins.columns[3].title = this.$t("swap.favorite");

      this.trade.columns[0].title = this.$t("swap.num");
      this.trade.columns[1].title = this.$t("swap.price");
      this.trade.columns[2].title = this.$t("swap.time");

      this.plate.columns[0].title = this.$t("swap.stall");
      this.plate.columns[1].title = this.$t("swap.price");
      this.plate.columns[2].title = this.$t("swap.num");
      this.plate.columns[3].title = this.$t("swap.total");

      this.currentOrder.columns[1].title = this.$t("swap.time");
      this.currentOrder.columns[2].title = this.$t("swap.symbol");
      this.currentOrder.columns[3].title = this.$t("swap.type");
      this.currentOrder.columns[4].title = this.$t("swap.direction");
      this.currentOrder.columns[5].title = this.$t("swap.price");
      this.currentOrder.columns[6].title = this.$t("swap.num");
      this.currentOrder.columns[7].title = this.$t("swap.traded");
      this.currentOrder.columns[8].title = this.$t("swap.dealamount");
      this.currentOrder.columns[9].title = this.$t("swap.action");

      this.historyOrder.columns[1].title = this.$t("swap.time");
      this.historyOrder.columns[2].title = this.$t("swap.symbol");
      this.historyOrder.columns[3].title = this.$t("swap.type");
      this.historyOrder.columns[4].title = this.$t("swap.direction");
      this.historyOrder.columns[5].title = this.$t("swap.price");
      this.historyOrder.columns[6].title = this.$t("swap.num");
      this.historyOrder.columns[7].title = this.$t("swap.done");
      this.historyOrder.columns[8].title = this.$t("swap.dealamount");
      this.historyOrder.columns[9].title = this.$t("swap.status");

      // window.tvWidget.options.time_frames[0].title = this.$t("swap.realtime");
    },
    getCNYRate() {
      this.$http
        .post(this.host + "/market/exchange-rate/usd/cny")
        .then((response) => {
          var resp = response.body;
          this.CNYRate = resp.data;
        });
    },
    getCoin(symbol) {
      return this.coins._map[symbol];
    },
    getKline() {
      var that = this;
      let config = {
        autosize: true,
        height: 500,
        fullscreen: false,
        symbol: that.symbol,
        interval: "60", // 默认K线周期
        timezone: "Asia/Shanghai",
        toolbar_bg: "#192330",
        container_id: "swap_kline_container",
        datafeed: that.datafeed,
        library_path:
          process.env.NODE_ENV === "production"
            ? "/assets/charting_library/"
            : "/src/assets/js/charting_library/",
        locale: "zh",
        debug: false,
        drawings_access: {
          type: "black",
          tools: [{ name: "Regression Trend" }],
        },
        disabled_features: [
          "header_resolutions",
          "timeframes_toolbar",
          "header_symbol_search",
          "header_chart_type",
          "header_compare",
          "header_undo_redo",
          "header_screenshot",
          "header_saveload",
          "use_localstorage_for_settings",
          "left_toolbar",
          "volume_force_overlay",
        ],
        enabled_features: [
          "hide_last_na_study_output",
          "move_logo_to_main_pane",
        ],
        custom_css_url: "bundles/common.css",
        supported_resolutions: ["1", "5", "15", "30", "60", "1D", "1W", "1M"],
        charts_storage_url: "http://saveload.tradingview.com",
        charts_storage_api_version: "1.1",
        client_id: "tradingview.com",
        user_id: "public_user_id",
        overrides: {
          "paneProperties.background": "#192330",
          "paneProperties.vertGridProperties.color": "rgba(0,0,0,.1)",
          "paneProperties.horzGridProperties.color": "rgba(0,0,0,.1)",
          //"scalesProperties.textColor" : "#AAA",
          "scalesProperties.textColor": "#61688A",
          "mainSeriesProperties.candleStyle.upColor": "#00b275",
          "mainSeriesProperties.candleStyle.downColor": "#f15057",
          "mainSeriesProperties.candleStyle.drawBorder": false,
          "mainSeriesProperties.candleStyle.wickUpColor": "#589065",
          "mainSeriesProperties.candleStyle.wickDownColor": "#AE4E54",
          "paneProperties.legendProperties.showLegend": false,

          "mainSeriesProperties.areaStyle.color1": "rgba(71, 78, 112, 0.5)",
          "mainSeriesProperties.areaStyle.color2": "rgba(71, 78, 112, 0.5)",
          "mainSeriesProperties.areaStyle.linecolor": "#9194a4",
          volumePaneSize: "small",
        },
        time_frames: [
          {
            text: "1min",
            resolution: "1",
            description: "realtime",
            title: that.$t("swap.realtime"),
          },
          { text: "1min", resolution: "1", description: "1min" },
          { text: "5min", resolution: "5", description: "5min" },
          { text: "15min", resolution: "15", description: "15min" },
          { text: "30min", resolution: "30", description: "30min" },
          {
            text: "1hour",
            resolution: "60",
            description: "1hour",
            title: "1hour",
          },
          {
            text: "4hour",
            resolution: "240",
            description: "4hour",
            title: "4hour",
          },
          {
            text: "1day",
            resolution: "1D",
            description: "1day",
            title: "1day",
          },
          {
            text: "1week",
            resolution: "1W",
            description: "1week",
            title: "1week",
          },
          { text: "1mon", resolution: "1M", description: "1mon" },
        ],
      };
      if (that.skin === "day") {
        config.toolbar_bg = "#fff";
        config.custom_css_url = "bundles/common_day.css";
        config.overrides["paneProperties.background"] = "#fff";
        config.overrides["mainSeriesProperties.candleStyle.upColor"] =
          "#a6d3a5";
        config.overrides["mainSeriesProperties.candleStyle.downColor"] =
          "#ffa5a6";
      }
      require(["@js/charting_library/charting_library.min.js"], function (tv) {
        var widget = (window.tvWidget = new TradingView.widget(config));
        widget.onChartReady(function () {
          widget.chart().executeActionById("drawingToolbarAction");
          widget
            .chart()
            .createStudy("Moving Average", false, false, [5], null, {
              "plot.color": "#EDEDED",
            });
          widget
            .chart()
            .createStudy("Moving Average", false, false, [10], null, {
              "plot.color": "#ffe000",
            });
          widget
            .chart()
            .createStudy("Moving Average", false, false, [30], null, {
              "plot.color": "#ce00ff",
            });
          widget
            .chart()
            .createStudy("Moving Average", false, false, [60], null, {
              "plot.color": "#00adff",
            });
          widget
            .createButton()
            .attr("title", "realtime")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(3);
              widget.setSymbol("", "1");
            })
            .append("<span>分时</span>");

          widget
            .createButton()
            .attr("title", "M1")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "1");
            })
            .append("<span>M1</span>");

          widget
            .createButton()
            .attr("title", "M5")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "5");
            })
            .append("<span>M5</span>");

          widget
            .createButton()
            .attr("title", "M15")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "15");
            })
            .append("<span>M15</span>");

          widget
            .createButton()
            .attr("title", "M30")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "30");
            })
            .append("<span>M30</span>");

          widget
            .createButton()
            .attr("title", "H1")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "60");
            })
            .append("<span>H1</span>")
            .addClass("selected");

          widget
            .createButton()
            .attr("title", "D1")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "1D");
            })
            .append("<span>D1</span>");

          widget
            .createButton()
            .attr("title", "W1")
            .on("click", function () {
              if ($(this).hasClass("selected")) return;
              $(this)
                .addClass("selected")
                .parent(".group")
                .siblings(".group")
                .find(".button.selected")
                .removeClass("selected");
              widget.chart().setChartType(1);
              widget.setSymbol("", "1W");
            })
            .append("<span>W1</span>");

          // widget
          //   .createButton()
          //   .attr("title", "M1")
          //   .on("click", function () {
          //     if ($(this).hasClass("selected")) return;
          //     $(this)
          //       .addClass("selected")
          //       .parent(".group")
          //       .siblings(".group")
          //       .find(".button.selected")
          //       .removeClass("selected");
          //     widget.chart().setChartType(1);
          //     widget.setSymbol("", "1M");
          //   })
          //   .append("<span>M1</span>");
        });
      });
    },
    getSymbol() {
      this.$http
        .post(this.host + this.api.swap.thumb, { type: this.currentCoin.type })
        .then((response) => {
          var resp = response.body;
          //先清空已有数据
          for (var i = 0; i < resp.length; i++) {
            var coin = resp[i];
            coin.base = resp[i].symbol.split("/")[1];
            this.coins[coin.base] = [];
            this.coins[coin.base + "2"] = [];
            this.coins._map = [];
          }
          for (var i = 0; i < resp.length; i++) {
            var coin = resp[i];
            coin.price = resp[i].close = resp[i].close.toFixed(
              this.baseCoinScale
            );
            coin.rose =
              resp[i].chg > 0
                ? "+" + (resp[i].chg * 100).toFixed(2) + "%"
                : (resp[i].chg * 100).toFixed(2) + "%";
            coin.coin = resp[i].symbol.split("/")[0];
            coin.base = resp[i].symbol.split("/")[1];
            coin.href = coin.id;
            coin.isFavor = false;
            if (!this.coins._map[coin.symbol]) {
              this.coins._map[coin.symbol] = [coin];
            } else {
              this.coins._map[coin.symbol].push(coin);
            }
            if (coin.zone == 0) {
              this.coins[coin.base].push(coin);
            } else {
              this.coins[coin.base + "2"].push(coin);
            }
            if (coin.id == this.currentCoin.id) {
              /**保留字段 */
              coin.perUsdt = this.currentCoin.perUsdt;
              coin.type = this.currentCoin.type;
              this.currentCoin = coin;
              this.form.limitPrice = this.form.limitPrice = parseFloat(
                coin.price
              );
            }
          }
          require(["../../assets/js/exchange.js"], function (e) {
            e.clickScTab();
          });
          this.startWebsock();
          this.changeBaseCion(this.basecion);
        });
    },
    getCoinInfo() {
      //获取精度
      this.$http
        .post(this.host + this.api.market.coinInfo, {
          unit: this.currentCoin.coin,
        })
        .then((response) => {
          var resp = response.body;
          if (resp != null) {
            this.coinInfo = resp;
          }
        });
    },
    getPlate(str = "") {
      //买卖盘
      var params = {};
      params["symbol"] = this.currentCoin.symbol;
      this.$http
        .post(this.host + this.api.swap.platemini, params)
        .then((response) => {
          this.plate.askRows = [];
          this.plate.bidRows = [];
          let resp = response.body;
          if (resp.ask && resp.ask.items) {
            for (var i = 0; i < resp.ask.items.length; i++) {
              if (i == 0) {
                resp.ask.items[i].totalAmount = resp.ask.items[i].amount;
              } else {
                resp.ask.items[i].totalAmount =
                  resp.ask.items[i - 1].totalAmount + resp.ask.items[i].amount;
              }
            }
            if (resp.ask.items.length >= this.plate.maxPostion) {
              for (var i = this.plate.maxPostion; i > 0; i--) {
                var ask = resp.ask.items[i - 1];
                ask.direction = "SELL";
                ask.position = i;
                this.plate.askRows.push(ask);
              }
              const rows = this.plate.askRows,
                len = rows.length,
                totle = rows[0].totalAmount;
              this.plate.askTotle = totle;
            } else {
              for (
                var i = this.plate.maxPostion;
                i > resp.ask.items.length;
                i--
              ) {
                var ask = { price: 0, amount: 0 };
                ask.direction = "SELL";
                ask.position = i;
                ask.totalAmount = ask.amount;
                this.plate.askRows.push(ask);
              }
              for (var i = resp.ask.items.length; i > 0; i--) {
                var ask = resp.ask.items[i - 1];
                ask.direction = "SELL";
                ask.position = i;
                this.plate.askRows.push(ask);
              }
              var askItemIndex =
                resp.ask.items.length - 1 > 0 ? resp.ask.items.length - 1 : 0;
              const rows = this.plate.askRows,
                len = rows.length,
                totle = rows[askItemIndex].totalAmount;
              this.plate.askTotle = totle;
            }
          }
          if (resp.bid && resp.bid.items) {
            for (var i = 0; i < resp.bid.items.length; i++) {
              if (i == 0)
                resp.bid.items[i].totalAmount = resp.bid.items[i].amount;
              else
                resp.bid.items[i].totalAmount =
                  resp.bid.items[i - 1].totalAmount + resp.bid.items[i].amount;
            }
            for (var i = 0; i < resp.bid.items.length; i++) {
              var bid = resp.bid.items[i];
              bid.direction = "BUY";
              bid.position = i + 1;
              this.plate.bidRows.push(bid);
              if (i == this.plate.maxPostion - 1) break;
            }
            if (resp.bid.items.length < this.plate.maxPostion) {
              for (
                var i = resp.bid.items.length;
                i < this.plate.maxPostion;
                i++
              ) {
                var bid = { price: 0, amount: 0 };
                bid.direction = "BUY";
                bid.position = i + 1;
                bid.totalAmount = 0;
                this.plate.bidRows.push(bid);
              }
              var bidItemIndex =
                resp.bid.items.length - 1 > 0 ? resp.bid.items.length - 1 : 0;
              const rows = this.plate.bidRows,
                len = rows.length,
                totle = rows[bidItemIndex].totalAmount;
              this.plate.bidTotle = totle;
            } else {
              const rows = this.plate.bidRows,
                len = rows.length,
                totle = rows[len - 1].totalAmount;
              this.plate.bidTotle = totle;
            }
          }
          if (str != "") {
            this.selectedPlate = str;
          }
        });
    },
    getPlateFull() {
      //深度图
      var params = {};
      params["symbol"] = this.currentCoin.symbol;
      this.$http
        .post(this.host + this.api.swap.platefull, params)
        .then((response) => {
          var resp = response.body;
          this.fullTrade = resp;
          resp.skin = this.skin;
          this.$refs.depthGraph.draw(resp);
        });
    },
    updatePlate(type, row) {
      //发现该方法未被使用
      if (type == "sell") {
        for (var i = 0; i < this.plate.askRows.length; i++) {
          if (
            row.price > this.plate.askRows[i].price &&
            i != 0 &&
            this.plate.askRows[i].price > 0
          ) {
            this.plate.askRows.splice(i, 0, row);
            this.plate.askRows.shift();
            break;
          } else if (
            i == this.plate.askRows.length - 1 &&
            (row.price < this.plate.askRows[i].price ||
              this.plate.askRows[i].price == 0)
          ) {
            this.plate.askRows.push(row);
            this.plate.askRows.shift();
            break;
          }
        }
      } else if (type == "buy") {
        for (var i = 0; i < this.plate.bidRows.length; i++) {
          if (row.price > this.plate.bidRows[i].price) {
            this.plate.bidRows.splice(i, 0, row);
            this.plate.bidRows.pop();
            break;
          }
        }
      }
    },
    startWebsock() {
      var that = this;
      var socket = io("http://localhost");
      this.datafeed = new Datafeeds.WebsockFeed(that.host+'/swap',this.currentCoin,socket,that.baseCoinScale);
      // this.getKline();
        that.getKline();
        //订阅价格变化消息
        socket.on("/topic/swap/thumb", function (msg) {
          var resp = JSON.parse(msg);
          var coins = that.getCoin(resp.symbol);
          if (coins != null && coins.length > 0) {
            coins.forEach((coin) => {
              // coin.price = resp.close.toFixed(2);
              coin.price = resp.close;
              coin.rose =
                resp.chg > 0
                  ? "+" + (resp.chg * 100).toFixed(2) + "%"
                  : (resp.chg * 100).toFixed(2) + "%";
              // coin.close = resp.close.toFixed(2);
              // coin.high = resp.high.toFixed(2);
              // coin.low = resp.low.toFixed(2);
              coin.close = resp.close;
              coin.high = resp.high;
              coin.low = resp.low;
              coin.turnover = parseInt(resp.volume);
              coin.volume = resp.volume;
              coin.usdRate = resp.usdRate;
            });
          }
        });
    },
    open_order() {
      this.showMarket = false;
    },
    close_order() {
      this.showMarket = true;
    },
    gohref(currentRow, oldCurrentRow) {
      this.$router.push({
        name: "SwapPairx",
        params: {
          pair: currentRow.href,
        },
      });
    },
    buyPlate(currentRow) {
      this.form.limitPrice = parseFloat(currentRow.price);
      if (this.form.type == 3) {
        this.form.triggerPrice = parseFloat(currentRow.price);
      }
    },
    sellPlate(currentRow) {
      this.form.limitPrice = parseFloat(currentRow.price);
      if (this.form.type == 3) {
        this.form.triggerPrice = parseFloat(currentRow.price);
      }
    },
    /**
     * 获取钱包信息
     */
    getWallet() {
      this.$http
        .post(this.host + this.api.uc.contractWallet + "USDT", {})
        .then((response) => {
          var resp = response.body;
          this.wallet.base = resp.data.balance;
          this.wallet.frozen = resp.data.frozenBalance;
        });
      this.$http
        .post(this.host + this.api.uc.wallet + "KICK", {})
        .then((response) => {
          var resp = response.body;
          this.wallet.kick = resp.data.balance;
        });
    },
    getCurrentPosition() {
      //当前持仓订单
      var params = {};
      params["contractCoinId"] = this.currentCoin.id;
      this.currentPosition.rows = [];
      var that = this;
      this.$http
        .post(this.host + this.api.swap.position, params)
        .then((response) => {
          var resp = response.body;
          if (resp.data.length > 0) {
            this.currentPosition.rows = resp.data.map((e) => {
              e.profit = 0.0;
              return e;
            });
          }
        });
    },
    getCurrentOrder() {
      //查询当前委托
      var params = {};
      params["pageNo"] = 0;
      params["pageSize"] = 100;
      params["contractCoinId"] = this.currentCoin.id;
      params["type"] = this.selectedType;
      this.currentOrder.rows = [];
      var that = this;
      this.$http
        .post(this.host + this.api.swap.current, params)
        .then((response) => {
          var resp = response.body;
          if (resp.content.length > 0) {
            this.currentOrder.rows = resp.content;
          }
        });
    },
    getHistoryOrder(pageNo) {
      //查询历史委托
      if (pageNo == undefined) {
        pageNo = this.historyOrder.page;
      } else {
        pageNo = pageNo - 1;
      }
      this.historyOrder.rows = []; //清空数据
      var params = {};
      params["pageNo"] = pageNo;
      params["pageSize"] = this.historyOrder.pageSize;
      params["contractCoinId"] = this.currentCoin.id;
      params["type"] = this.selectedType;
      var that = this;
      this.$http
        .post(this.host + this.api.swap.history, params)
        .then((response) => {
          var resp = response.body;
          let rows = [];
          if (resp.content.length > 0) {
            this.historyOrder.total = resp.totalElements;
            this.historyOrder.page = resp.number;
            this.historyOrder.rows = resp.content;
          }
        });
    },
    cancel(index) {
      var order = this.currentOrder.rows[index];
      this.$Modal.confirm({
        title: this.$t("swap.tip"),
        content: this.$t("swap.undotip"),
        onOk: () => {
          this.$http
            .post(this.host + this.api.swap.entrustCancel + "/" + order.id, {})
            .then((response) => {
              var resp = response.body;
              if (resp.code == 0) {
                this.refreshAccount();
                this.$Notice.success({
                  title: this.$t("swap.tip"),
                  desc: this.$t("swap.cancelsuccess"),
                });
              } else {
                this.$Notice.error({
                  title: this.$t("swap.tip"),
                  desc: resp.message,
                });
              }
            });
        },
      });
    },
    refreshAccount: function () {
      this.getCurrentOrder();
      this.getHistoryOrder();
      this.getCurrentPosition();
      this.getWallet();
    },
    timeFormat: function (tick) {
      return moment(tick).format("HH:mm:ss");
    },
    dateFormat: function (tick) {
      return moment(tick).format("YYYY-MM-DD HH:mm:ss");
    },
    keyEvent(event) {
      var re1 = new RegExp(
        "([0-9]+.[0-9]{" + this.baseCoinScale + "})[0-9]*",
        ""
      );
      this.form.limitPrice = this.form.limitPrice.toString().replace(re1, "$1");
      this.form.limitPrice = this.form.limitPrice.toString().replace(re1, "$1");
      this.form.marketAmount = this.form.marketAmount
        .toString()
        .replace(re1, "$1");

      var re2 = new RegExp("([0-9]+.[0-9]{" + this.coinScale + "})[0-9]*", "");
      this.form.limitAmount = this.form.limitAmount
        .toString()
        .replace(re2, "$1");
      this.form.limitAmount = this.form.limitAmount
        .toString()
        .replace(re2, "$1");
      this.form.marketAmount = this.form.marketAmount
        .toString()
        .replace(re2, "$1");
    },
    showMarginModeModal() {
      this.marginModeModal = true;
      console.log("hello");
    },
    changeMarginMode(val) {
      this.form.patterns = val;
    },
    showAdjustLeverage(type) {
      this.form.leverageModal = true;
      this.form.leverageAdjustVal = this.form.leverage;
    },
    showAdjustTime() {
      this.form.timeModal = true;
      this.form.timeAdjustVal = this.form.time;
      if (this.form.time == 10) {
        this.form.timeAdjustVal2 = 1;
      } else if (this.form.time == 30) {
        this.form.timeAdjustVal2 = 2;
      } else if (this.form.time == 60) {
        this.form.timeAdjustVal2 = 3;
      } else if (this.form.time == 120) {
        this.form.timeAdjustVal2 = 4;
      } else if (this.form.time == 360) {
        this.form.timeAdjustVal2 = 5;
      } else if (this.form.time == 3000) {
        this.form.timeAdjustVal2 = 6;
      }
    },
    adjustTime() {
      this.form.time = this.form.timeAdjustVal;
      this.form.timeModal = false;
    },
    adjustLeverage() {
      if (!this.isLogin) {
        this.$Notice.error({
          title: this.$t("common.tip"),
          desc: this.$t("common.logintip"),
        });
        return;
      }
      this.$http
        .post(this.host + this.api.swap.leverage, {
          symbol: this.currentCoin.symbol,
          leverage: this.form.leverageAdjustVal,
        })
        .then((response) => {
          var resp = response.body;
          if (resp.code == 0) {
            this.form.leverage = this.form.leverageAdjustVal;
            this.form.leverageModal = false;
          } else {
            this.$Notice.error({
              title: this.$t("swap.tip"),
              desc: this.$t("swap.modifyLeverageWrong"),
            });
            this.form.leverageModal = false;
          }
        });
    },
  },
};
</script>
<style scoped>
.footer, .layout  {
  display: none;
}
#tawkchat-minified-box {
  display: none;
}
</style>
<style scoped lang="scss">
@import url(../../assets/css/swap.css);
$night-bg: #0b1520;
$night-headerbg: #27313e;
$night-contentbg: #192330;
:root { $night-color: #fff; }
::v-deep.ivu-slider {
  .ivu-slider-wrap {
    height: 10px;
    border-radius: 0px;
    left: -6px;
    .ivu-slider-bar {
      height: 10px;
    }
    .ivu-slider-stop {
      height: 10px;
      width: 10px;
      top: 0px;
      border-radius: 0px;
      margin-left: 5px;
    }
    .ivu-slider-button-wrap {
      .ivu-tooltip {
        .ivu-tooltip-rel {
          .ivu-slider-button {
            width: 24px;
            height: 24px;
            margin-top: -4px;
          }
        }
      }
    }
    .ivu-slider-marks {
      .ivu-slider-marks-item {
        margin-left: 4px;
      }
    }
  }
}
::v-deep.v-transfer-dom {
  .ivu-modal-wrap {
    .ivu-modal {
      .ivu-modal-content {
        .ivu-modal-body {
          width: 90%;
          margin: 0 auto;
        }
      }
    }
  }
}
.account-item {
  width: 90%;
  margin: 0 auto;
}
.account-item > div:nth-child(1) {
  float: left;
  line-height: 38px;
  font-size: 12px;
  color: #b0b8db;
}
.account-item > div:nth-child(2) {
  float: right;
  line-height: 38px;
  color: #fff;
}
.ivu-btn-small {
  padding: none;
}
.ivu-radio-group-button .ivu-radio-wrapper-checked:first-child {
  color: #000;
  border-color: rgb(240, 172, 25);
  background: rgb(240, 172, 25);
  box-shadow: none;
}
.ivu-radio-group-button .ivu-radio-wrapper-checked:first-child:hover {
  color: #fff;
}
.ivu-radio-group-button .ivu-radio-wrapper:hover {
  color: #fff;
}
.ivu-radio-group-button .ivu-radio-wrapper {
  height: 25px;
  line-height: 25px;
  color: #fff;
  background: #000;
  border: 1px solid #fff;
}
.ivu-radio-group-button .ivu-radio-wrapper-checked {
  color: #000;
  border-color: rgb(240, 172, 25);
  background: rgb(240, 172, 25);
  box-shadow: none;
}
::v-deep.v-transfer-dom {
  .ivu-modal-wrap {
    .ivu-modal {
      top: 30%;
    }
  }
}
.operate_details1 span {
  color: #41b371;
  font-size: 12px;
}
.operate_details2 span {
  color: #d74e5a;
  font-size: 12px;
}
.operate_left .operate_btn {
  background: #41b371;
}
.operate_right .operate_btn {
  background: #d74e5a;
}
.operate_btn {
  padding: 3px 15px;
  border-radius: 4px;
  margin-top: 10px;
}
.swap .order {
  margin-bottom: 5px !important;
}
::v-deep.ivu-input-group .ivu-input {
  color: #f0ac19;
}

.order-type-btn-wrap {
  height: 30px;
  padding: 0 16px;
  width: 100%;
  background: #192330;
  button {
    width: auto;
    border-radius: 5px;
    height: auto;
    background: #192330;
    border: 1px solid #fff;
    padding: 3px 15px;
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
.swap {
  color: #fff;
  background-color: #0b1520;
  padding-top: 0px;
  .main {
    width: 100%;
    margin-left:0;
    display: flex;
    margin-top: 0;
    .left {
      border-radius: 0px;
      margin-right: 5px;
      overflow: hidden;
      .handlers {
        font-size: 0;
        padding: 10px 20px;
        background-color: #192330;
        border-top-left-radius: 0px;
        border-top-right-radius: 0px;
        .handler {
          display: inline-block;
          margin-right: 10px;
          width: 20px;
          height: 20px;
          background-size: 100% 100%;
          cursor: pointer;
          &.handler-all {
            background-image: url("../../assets/images/exchange/plate_all.png");
            &.active {
              background-image: url("../../assets/images/exchange/plate_all_active.png");
            }
          }
          &.handler-green {
            background-image: url("../../assets/images/exchange/plate_green.png");
            &.active {
              background-image: url("../../assets/images/exchange/plate_green_active.png");
            }
          }
          &.handler-red {
            background-image: url("../../assets/images/exchange/plate_red.png");
            &.active {
              background-image: url("../../assets/images/exchange/plate_red_active.png");
            }
          }
        }
      }
      .plate-nowprice {
        text-align: center;
        background-color: #27313e;
        line-height: 1;
        display: flex;
        align-items: center;
        height: 40px;
        justify-content: center;
        font-size: 14px;
        font-weight: 500;
        .price {
          font-size: 18px;
          margin-right: 10px;
        }
        .price-cny {
          font-size: 12px;
          margin-left: 10px;
          font-weight: 400;
          color: rgba(219, 222, 246, 0.3);
        }
      }
    }
    .center {
      z-index: 10000000000;
      flex: 0 0 100%;
      margin-right: 5px;
      .imgtable {
        height: 446px;
        position: relative;
        overflow: hidden;
        .handler {
          position: absolute;
          top: 10px;
          right: 40px;
          z-index: 1000;
          > span {
            background-color: #2c3b59;
            color: #c7cce6;
            padding: 4px 8px;
            cursor: pointer;
            font-size: 13px;
            opacity: 0.5;
            &.active {
              opacity: 1;
            }
          }
        }
      }
      .trade_wrap {
        .trade_menu {
          position: relative;
          background-color: #192330;
          border-top-left-radius: 0px;
          border-top-right-radius: 0px;
          border-bottom: 1px solid #27313e;
          font-size: 0;
          height: 40px;
          line-height: 40px;
          > span {
            font-size: 16px;
            padding: 11px 20px;
            cursor: pointer;
            &.active {
              color: #fff;
              border-bottom: 2px solid #f0a70a;
            }
            &:first-child {
              border-top-left-radius: 0px;
            }
          }
          .fee-wrap {
            position: absolute;
            top: 0;
            right: 20px;
            font-size: 12px;
            > span {
              margin-right: 10px;
              color: #7c7f82;
            }
            > a {
              vertical-align: middle;
            }
          }
        }
        .trade_panel {
          position: relative;
          .mask {
            position: absolute;
            width: 100%;
            height: 100%;
            display: flex;
            background-color: rgba(0, 52, 77, 0.8);
            justify-content: center;
            align-items: center;
            z-index: 100;
            font-size: 24px;
            border-radius: 0px;
            // color: #bcd7e6;
          }
          .publish-mask {
            position: absolute;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 37, 64, 0.93);
            justify-content: center;
            align-items: center;
            z-index: 101;
            font-size: 24px;
            border-radius: 0px;
            // color: #bcd7e6;
          }
        }
        .trade_panel .panel .hd {
          border-bottom: none;
          b {
            color: #fff;
          }
          a {
            color: #f0a70a;
          }
        }
      }
    }
    .right {
      flex: 0 0 23%;
      .coin-menu {
        overflow: hidden;
        height: 512px;
        background-color: #192330;
        border-radius: 0px;
      }
    }
  }
  .symbol {
    display: flex;
    justify-content: space-between;
    padding: 15px 30px;
    margin-bottom: 5px;
    align-items: center;
    background-color: #192330;
    line-height: 1;
    border-radius: 0px;
    .item {
      .price-cny {
        font-size: 12px;
        color: #546886;
      }
      .coin {
        font-size: 20px;
      }
      .text {
        width: 100%;
        display: block;
        font-size: 12px;
        color: #999;
        margin-bottom: 5px;
      }
      .num {
        font-size: 12px;
        color: #fff;
      }
      > img {
        display: block;
        width: 18px;
        height: 18px;
        cursor: pointer;
      }
    }
  }
  .order {
    min-height: 227px;
    margin-bottom: 10px;
    overflow: hidden;
    .order-handler {
      // background-color: #192330;
      background-color: #192330;
      font-size: 0;
      // line-height: 38px;
      > span {
        padding: 0 20px;
        font-size: 14px;
        display: inline-block;
        color: #fff;
        cursor: pointer;
        line-height: 30px;
        background-color: #192330;
        &.active {
          color: #f0a70a;
          background-color: #27313e;
          border-bottom: 2px solid #f0a70a;
        }
        &:first-child {
          border-top-left-radius: 0px;
        }
        &:last-child {
          border-top-right-radius: 0px;
        }
      }
    }
  }
}
.exchange.day {
  color: #333;
  background-color: #fff;
  .main {
    .left {
      background-color: #fff;
      box-shadow: 0 0 2px #ccc;
      .handlers {
        background-color: #fff;
      }
      .plate-nowprice {
        background-color: #fff;
        border-top: 1px solid #f0f0f0;
        border-bottom: 1px solid #f0f0f0;
      }
    }
    .imgtable {
      border-radius: 6px;
      box-shadow: 0 0 2px #ccc;
      .handler {
        > span {
          border: 1px solid #333;
        }
      }
    }
    .trade_wrap {
      box-shadow: 0 0 2px #ccc;
      .trade_menu {
        background-color: #fafafa;
        > span {
          background-color: #fafafa;
          border-right: 1px solid #f0f0f0;
          &.active {
            background-color: #fff;
            color: #f0a70a;
          }
          &:last-child {
            border-top-right-radius: 6px;
          }
        }
        .ivu-icon {
          color: #333 !important;
        }
      }
      .trade_panel {
        box-shadow: 0 0 2px #ccc;
        .mask {
          background-color: rgba(0, 52, 77, 0.8);
          color: #fff;
        }
      }
      .trade_panel .panel .hd {
        border-bottom: none;
        b {
          color: #333;
        }
        a {
          color: #f0a70a;
        }
      }
    }
    .right {
      .coin-menu {
        background-color: #fff;
        box-shadow: 0 2px 2px #ccc;
      }
      .trade-wrap {
        box-shadow: 0 0 2px #ccc;
        border-radius: 6px;
      }
      // .ivu-table-wrapper{
      //         box-shadow:0 0 2px #ccc;
      //       }
    }
  }
  .symbol {
    background-color: #fff;
    box-shadow: 0 0 2px #ccc;
    .item {
      .text {
        color: #999;
      }
      .num {
        color: #333;
      }
    }
  }
  .order {
    box-shadow: 0 2px 2px #ccc;
    min-height: 227px;
    .order-handler {
      margin-right: -2px;
      background-color: #fff;
      > span {
        color: #333;
        background-color: #fafafa;
        box-shadow: 0 0 2px #ccc;
        &.active {
          color: #f0a70a;
          background-color: #fff;
        }
      }
    }
  }
}
#swap_kline_container {
  background: #192330;
}

.coin-info {
  color: #8f9ca5;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 5;
  overflow: hidden;
  padding-top: 15px;
}
</style>