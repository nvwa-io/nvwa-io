<template>
  <div style="height: 40px;">
    <div class="header">

      <el-row :gutter="32">
        <el-col :xs="10" :sm="10" :lg="10">
          <div class="logo">
            <router-link to="/home">
              <img src="@/assets/images/app-white.png" alt="">
              <span>女娲</span>
            </router-link>
          </div>
        </el-col>
        <el-col :xs="14" :sm="14" :lg="14">
          <div class="website">
            <!--<span v-if="!isLogin">-->
            <span v-if="isLogin">
              <a @click="handleShowProjects">
                进入项目
              </a>
              <router-link to="/home/project">
                管理项目
              </router-link>
              <router-link to="/home/audit">
                <div style="display: inline-block;">工单
                  <el-badge v-if="auditNum > 0" :value="auditNum" class="mark audit-badge"/>
                </div>
              </router-link>
              <router-link to="/admin/index">
                管理端
              </router-link>
            </span>
            <!--<span v-else>-->
            <span v-else>
              <a @click="handleLogin">
                登录
              </a>
              <a @click="handleRegister">
                注册
              </a>
            </span>
            <span style="min-height:8px;width: 1px;background: rgba(78,106,140,0.91); display: inline-block;margin: 0 5px;"/>
            <a href="http://nvwa-io.com" target="_blank">
              女娲官网
            </a>
          </div>
        </el-col>
      </el-row>
    </div>

    <el-dialog
      :visible.sync="dialogSelectProject"
      :show-close="true"
      class="login-container"
      title="进入项目"
      style="margin-top: 0; padding-top: 0;">
      <div style="min-height: 180px;margin-top: -30px;">
        <a
          v-for="p in projects"
          :key="p.project.id"
          class="project-item"
          @click="handleSelectProject(p.project)">
          {{ p.project.name }}
        </a>

        <p v-if="projects.length===0" class="no-project-info">暂无项目, <router-link to="/project"> 创建一个项目</router-link></p>
      </div>
    </el-dialog>

    <el-dialog
      :visible.sync="dialogLoginVisible"
      :show-close="true"
      width="400px"
      class="login-container"
      title="登录"
      style="margin-top: 0px; padding-top: 0px;">
      <div style="margin-top: -20px;">
        <el-menu :default-active="activeLoginIndex" mode="horizontal" @select="handleLoginTabSelect">
          <el-menu-item index="ldap">Ldap 登录</el-menu-item>
          <el-menu-item index="default">默认登录</el-menu-item>
        </el-menu>
        <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
          <el-form-item prop="username">
            <span class="svg-container">
              <svg-icon icon-class="user" />
            </span>
            <el-input
              v-model="loginForm.username"
              :placeholder="$t('login.username')"
              name="username"
              type="text"
              auto-complete="on"
            />
          </el-form-item>

          <el-form-item prop="password">
            <span class="svg-container">
              <svg-icon icon-class="password" />
            </span>
            <el-input
              :type="passwordType"
              v-model="loginForm.password"
              :placeholder="$t('login.password')"
              name="password"
              auto-complete="on"
              @keyup.enter.native="handleLoginSubmit" />
            <span class="show-pwd" @click="showPwd">
              <svg-icon icon-class="eye" />
            </span>
          </el-form-item>
          <div style="padding-top: 20px; text-align: right;">
            <el-button @click="dialogLoginVisible = false">取消</el-button>
            <el-button :loading="loading" type="primary" style="width: 180px;" @click="handleLoginSubmit">登录</el-button>
          </div>
        </el-form>

      <!--<div slot="footer" class="dialog-footer">-->
      <!--<el-button type="primary" @click="handleLoginSubmit">登录</el-button>-->
      <!--</div>-->
      </div>

    </el-dialog>

    <el-dialog
      :visible.sync="dialogRegisterVisible"
      :show-close="true"
      width="400px"
      class="login-container"
      title="注册"
      style="margin-top: 0px; padding-top: 0px;">
      <el-form ref="registerForm" :model="loginForm" :rules="loginRules" style="padding-top: 0px" class="login-form" auto-complete="off" label-position="left">
        <el-form-item prop="username">
          <span class="svg-container">
            <svg-icon icon-class="user" />
          </span>
          <el-input
            v-model="loginForm.username"
            :placeholder="$t('login.username')"
            name="username"
            type="text"
            auto-complete="on"
          />
        </el-form-item>

        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :type="passwordType"
            v-model="loginForm.password"
            :placeholder="$t('login.password')"
            name="password"
            auto-complete="on"
            @keyup.enter.native="handleLogin" />
          <span class="show-pwd" @click="showPwd">
            <svg-icon icon-class="eye" />
          </span>
        </el-form-item>
        <el-form-item prop="username">
          <span class="svg-container">
            <svg-icon icon-class="user" />
          </span>
          <el-input
            v-model="loginForm.username"
            :placeholder="$t('login.username')"
            name="username"
            type="text"
            auto-complete="on"
          />
        </el-form-item>
        <div style="padding-top: 20px; text-align: right;">
          <el-button @click="dialogRegisterVisible = false">取消</el-button>
          <el-button type="success" style="width: 180px;" @click="handleLoginSubmit">注册</el-button>
        </div>
      </el-form>

    <!--<div slot="footer" class="dialog-footer">-->
    <!--<el-button type="primary" @click="handleLoginSubmit">登录</el-button>-->
    <!--</div>-->

    </el-dialog>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { isValidUsername } from '@/utils/validate'
import { getToken } from '@/utils/auth'
import gc from '@/utils/globalCache'
import { Message } from 'element-ui'
import apiProject from '@/api/project'

export default {
  data() {
    const validateUsername = (rule, value, callback) => {
      if (!isValidUsername(value)) {
        callback(new Error('用户名格式错误'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('密码最少 6 个字符'))
      } else {
        callback()
      }
    }
    return {
      dialogLoginVisible: false,
      activeLoginIndex: 'ldap',
      dialogRegisterVisible: false,
      dialogSelectProject: false,

      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      passwordType: 'password',
      loading: false,
      showDialog: false,
      redirect: undefined,

      isLogin: getToken(),

      projects: []
    }
  },
  computed: {
    ...mapGetters([
      'auditNum'
    ])
  },
  created() {
    // reset store curProject from localstorage when page reload
    if (this.$store.state.project.curProject.id === 0) {
      if (!gc.getLatestProject()) {
        this.$router.push('/home')
        return
      }

      this.$store.dispatch('SetCurProject', gc.getLatestProject()).then(() => {
      })
    }

    // refresh user's audit num
    this.$store.dispatch('GetWaitAuditNum')
  },

  methods: {
    handleLogin() {
      this.dialogLoginVisible = true
    },

    handleLoginTabSelect(key, keyPath) {
      console.log(key, keyPath)
      this.activeLoginIndex = key
    },

    // submit login request
    handleLoginSubmit() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true

          console.log(this.loginForm)
          this.$store.dispatch('Login', {
            username: this.loginForm.username,
            password: this.loginForm.password,
            loginType: this.activeLoginIndex
          }).then(response => {
            this.loading = false
            this.$router.push({ path: this.redirect || '/' })
          }).catch((error) => {
            this.loading = false
            console.log(error)
            Message.error(error)
          })
        } else {
          console.log('Login form fields validate failed.')
          return false
        }
      })
    },
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
    },
    handleRegister() {
      this.dialogRegisterVisible = true
    },

    handleShowProjects() {
      this.dialogSelectProject = true
      apiProject.list().then(response => {
        this.projects = response.data.list
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },

    handleSelectProject(project) {
      const tmpProject = this.$store.state.project.curProject
      gc.setLatestProject(project)
      this.dialogSelectProject = false
      this.$store.dispatch('SetCurProject', project)

      if (!tmpProject || this.$route.path.indexOf('/home') === 0) {
        this.$router.push({ path: '/deploy/index' })
      } else {
        if (tmpProject.id === project.id) {
          return
        }

        // select another project, refresh page
        this.$router.go(0)
      }
    }

  }
}

</script>
<style rel="stylesheet/scss" lang="scss">
  /* 修复input 背景不协调 和光标变色 */
  /* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

  $bg:#fff;
  $light_gray:#fff;
  $cursor: #2a394b;

  @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .login-container .el-input input{
      color: $cursor;
      &::first-line {
        color: #2a394b;
      }
    }
  }

  /* reset element-ui css */
  .login-container {
    .el-input {
      display: inline-block;
      height: 47px;
      width: 85%;
      input {
        background: transparent;
        border: 0px;
        -webkit-appearance: none;
        border-radius: 0px;
        padding: 12px 5px 12px 15px;
        height: 47px;

        &:-webkit-autofill {
          -webkit-box-shadow: 0 0 0px 1000px $bg inset !important;
          -webkit-text-fill-color: $cursor !important;
        }
      }

    }
    .el-form-item {
      border: 1px solid rgba(16, 92, 188, 0.1);
      border-radius: 2px;
      color: #2a394b;
    }
  }
</style>

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

  .project-item {
    display: inline-block;
    margin: 15px 5px 15px 0;
    border: 1px solid rgba(100, 145, 196, 0.09);
    color: #1f2d3d;
    padding: 10px 25px;
    border-radius: 2px;
    &:hover {
      background-color: #1f2d3d;
      color: white;
    }
  }

  .no-project-info {
    text-align: center;
    color: rgba(42,57,75,0.45);
    margin-top: 100px;
    a {
      color: #2c3f56;
      font-size: 14px;
      margin: auto 10px;
    }
    a:hover {
      color: #141b25;
    }
  }

  $dark_gray:#889aa4;
  $light_gray:#eee;

  .login-container {
    /*position: fixed;*/
    height: 100%;
    width: 100%;
    .login-form {
      /*position: absolute;*/
      /*left: 0;*/
      /*right: 0;*/
      width: 100%;
      max-width: 100%;
      padding: 35px 0;
      margin: 0 auto;
    }
    .tips {
      font-size: 14px;
      margin-bottom: 10px;
      span {
        &:first-of-type {
          margin-right: 16px;
        }
      }
    }
    .svg-container {
      padding: 6px 5px 6px 15px;
      color: $dark_gray;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }
    .title-container {
      position: relative;
      .title {
        font-size: 26px;
        color: $light_gray;
        margin: 0px auto 40px auto;
        text-align: center;
        font-weight: bold;
      }
      .set-language {
        color: #fff;
        position: absolute;
        top: 5px;
        right: 0px;
      }
    }
    .show-pwd {
      position: absolute;
      right: 10px;
      top: 7px;
      font-size: 16px;
      color: $dark_gray;
      cursor: pointer;
      user-select: none;
    }
  }

  .el-menu--horizontal > .el-menu-item {
    font-size: 15px;
  }
  .el-menu--horizontal {
    border-bottom: solid 0 #fff;
  }
</style>
