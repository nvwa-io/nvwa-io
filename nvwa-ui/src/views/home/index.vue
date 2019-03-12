<template>
  <div>

    <el-carousel :interval="5000" indicator-position="outside">
      <el-carousel-item v-for="item in banners" :key="item.id">
        <img :src="item.src" style="width: 100%;" alt="">
        <div style="position: absolute; top: 0;z-index: 19;width:100%">
          <div style="width: 100%; margin: 100px auto 0;text-align: center;">
            <h3 style="font-size: 36px;color: #fff">
              女娲 - DevOps 部署系统
            </h3>
            <p style="color: #fff;" v-html="item.text"/>
          </div>
        </div>
      </el-carousel-item>

    </el-carousel>
    <div class="dashboard-editor-container">

      <panel-group @handleSetLineChartData="handleSetLineChartData"/>

      <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
        <line-chart :chart-data="lineChartData"/>
      </el-row>

      <el-row :gutter="32">
        <el-col :xs="24" :sm="24" :lg="8">
          <div class="chart-wrapper">
            <raddar-chart/>
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :lg="8">
          <div class="chart-wrapper">
            <pie-chart/>
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :lg="8">
          <div class="chart-wrapper">
            <bar-chart/>
          </div>
        </el-col>
      </el-row>

      <!--<el-row :gutter="8">-->
      <!--<el-col-->
      <!--:xs="{span: 24}"-->
      <!--:sm="{span: 24}"-->
      <!--:md="{span: 24}"-->
      <!--:lg="{span: 12}"-->
      <!--:xl="{span: 12}"-->
      <!--style="padding-right:8px;margin-bottom:30px;">-->
      <!--<transaction-table/>-->
      <!--</el-col>-->
      <!--<el-col-->
      <!--:xs="{span: 24}"-->
      <!--:sm="{span: 12}"-->
      <!--:md="{span: 12}"-->
      <!--:lg="{span: 6}"-->
      <!--:xl="{span: 6}"-->
      <!--style="margin-bottom:30px;">-->
      <!--<todo-list/>-->
      <!--</el-col>-->
      <!--<el-col-->
      <!--:xs="{span: 24}"-->
      <!--:sm="{span: 12}"-->
      <!--:md="{span: 12}"-->
      <!--:lg="{span: 6}"-->
      <!--:xl="{span: 6}"-->
      <!--style="margin-bottom:30px;">-->
      <!--<box-card/>-->
      <!--</el-col>-->
      <!--</el-row>-->
    <!---->
    </div>

  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import RaddarChart from './components/RaddarChart'
import PieChart from './components/PieChart'
import BarChart from './components/BarChart'
import TransactionTable from './components/TransactionTable'
import TodoList from './components/TodoList'
import BoxCard from './components/BoxCard'

const lineChartData = {
  newVisitis: {
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}

export default {
  name: 'Home',
  components: {
    PanelGroup,
    LineChart,
    RaddarChart,
    PieChart,
    BarChart,
    TransactionTable,
    TodoList,
    BoxCard
  },
  data() {
    return {
      lineChartData: lineChartData.newVisitis,

      banners: [
        {
          id: 0,
          src: require('@/assets/images/home-bg.jpg'),
          text: '女娲（<a href="http://nvwa-io.com" target="_blank">nvwa-io.com</a>）是一款基于 DevOps 理念的开源部署系统。'
        },
        {
          id: 1,
          src: require('@/assets/images/home-bg01.jpg'),
          text: '女娲以项目/应用（服务）为中心，与构建系统（Jenkins，可选）结合，提供完善的持续集成（CI）、持续交付（CD）解决方案。'
        }
      ]
    }
  },
  created: function() {
    console.log('=======created')
  },
  methods: {
    handleSetLineChartData(type) {
      this.lineChartData = lineChartData[type]
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .dashboard-editor-container {
    padding: 32px;
    background-color: rgb(240, 242, 245);
    .chart-wrapper {
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

</style>
