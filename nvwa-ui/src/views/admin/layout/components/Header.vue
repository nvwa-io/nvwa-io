<template>
  <div style="height: 40px">
    <div class="header">

      <el-row :gutter="32">
        <el-col :xs="10" :sm="10" :lg="10">
          <a href="/home">
            <div class="logo">
              <router-link to="/home">
                <img src="@/assets/images/app-white.png" alt="">
                <span>女娲</span>
              </router-link>
            </div>
          </a>
        </el-col>
        <el-col :xs="14" :sm="14" :lg="14">
          <div class="website">
            <router-link to="/home/index">
              用户端
            </router-link>
            <span style="min-height:8px;width: 1px;background: rgba(78,106,140,0.91); display: inline-block;margin: 0 5px;"/>
            <a href="http://nvwa-io.com" target="_blank">
              女娲官网
            </a>
        </div></el-col>
      </el-row>
    </div>
  </div>
</template>

<script>

</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .header {
    height: 40px;
    width: 100%;
    background: #1f2d3d;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
  }

  .logo {
    padding: 10px 15px;
    img {
      width: 20px;
      margin-right: 10px;
      float: left;
    }
    span {
      float: left;
      color: white;
    }
  }

  .website {
    text-align: right;
    padding: 10px;
    a {
      color: white;
      font-size: 14px;
      margin: auto 10px;
    }
    a:hover {
      color: #7d8899;
    }
  }

</style>
