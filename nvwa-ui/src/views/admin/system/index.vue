<template>
  <div>
    <div class="app-container">

      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="config-deploy">部署</el-menu-item>
        <el-menu-item index="config-build">构建</el-menu-item>
        <el-menu-item index="config-version-pkg">版本包</el-menu-item>
        <el-menu-item index="config-git">Git 认证</el-menu-item>
        <el-menu-item index="config-notify">通知</el-menu-item>
        <el-menu-item index="config-user">登录/注册</el-menu-item>
        <el-menu-item index="config-project">项目权限</el-menu-item>
      </el-menu>

      <br>

      <el-form ref="form" :model="form" label-width="160px">

        <el-row :gutter="30">
          <el-col :xs="18" :sm="18" :lg="18">

            <div v-show="activeIndex === 'config-deploy'">
              <el-form-item label="部署根路径">
                <el-input v-model="form.deployRootPath"/>
              </el-form-item>
              <el-form-item label="允许自定义部署路径">
                <el-switch v-model="form.customDeployPath"/>
              </el-form-item>
              <el-form-item label="默认部署用户">
                <el-input v-model="form.deployUser"/>
              </el-form-item>
              <el-form-item label="允许自定义部署用户">
                <el-switch v-model="form.customDeployUser"/>
              </el-form-item>
            </div>

            <div v-show="activeIndex === 'config-build'">
              <el-form-item label="本地版本包根路径">
                <el-input v-model="form.pkgRootPath"/>
              </el-form-item>
              <el-form-item label="本地代码仓库根路径">
                <el-input v-model="form.repoRootPath"/>
              </el-form-item>
              <el-form-item label="本地构建空间根路径">
                <el-input v-model="form.buildRootPath"/>
              </el-form-item>
              <el-form-item label="使用 Jenkins 构建">
                <el-switch v-model="form.useJenkins"/>
              </el-form-item>
              <el-form-item label="Jenkins 服务地址">
                <el-input v-model="form.jenkinsUrl"/>
              </el-form-item>
              <el-form-item label="Jenkins 用户">
                <el-input v-model="form.jenkinsUser"/>
              </el-form-item>
              <el-form-item label="Jenkins 密码">
                <el-input v-model="form.jenkinsPassword"/>
              </el-form-item>
              <el-form-item label="Jenkins 创建应用模板">
                <el-input v-model="form.jenkinsTemplate" type="textarea"/>
              </el-form-item>
            </div>

            <div v-show="activeIndex === 'config-git'">
              <el-form-item label="Git 认证方式">
                <el-radio-group v-model="form.gitCIAuthType">
                  <el-radio label="账号密码"/>
                  <el-radio label="Token"/>
                  <el-radio label="SSH 免密"/>
                </el-radio-group>
              </el-form-item>

              <el-form-item label="Git 用户名">
                <el-input v-model="form.gitCIUser"/>
              </el-form-item>

              <el-form-item label="Git 密码">
                <el-input v-model="form.gitCIPassword"/>
              </el-form-item>

              <el-form-item label="Git Token">
                <el-input v-model="form.gitCIToken"/>
              </el-form-item>
            </div>

            <div v-show="activeIndex === 'config-version-pkg'">
              <el-form-item label="服务器版本包根路径">
                <el-input v-model="form.pkgRootPath"/>
              </el-form-item>

              <el-form-item label="版本包仓库存储">
                <el-select v-model="form.pkgStorageType" placeholder="请选择存储方式">
                  <el-option label="主机本地" value="shanghai"/>
                  <el-option label="阿里云 OSS" value="beijing"/>
                </el-select>
              </el-form-item>

              <el-form-item label="版本包保留数量">
                <el-input-number v-model="form.pkgLimit" :min="1" :max="10" label="描述文字"/>
              </el-form-item>

              <el-form-item label="OSS Endpoint">
                <el-input v-model="form.pkgStorageConfig.oss.endpoint"/>
              </el-form-item>

              <el-form-item label="OSS AccessKey">
                <el-input v-model="form.pkgStorageConfig.oss.accessKey"/>
              </el-form-item>

              <el-form-item label="OSS AccessSecret">
                <el-input v-model="form.pkgStorageConfig.oss.accessSecret"/>
              </el-form-item>

              <el-form-item label="OSS Bucket">
                <el-input v-model="form.pkgStorageConfig.oss.bucket"/>
              </el-form-item>
            </div>

            <div v-show="activeIndex === 'config-notify'">
              <el-form-item label="通知类型">
                <el-select
                  v-model="selectedNotifyTypes"
                  multiple
                  filterable
                  allow-create
                  default-first-option
                  placeholder="请选择">
                  <el-option
                    v-for="item in selectNotifyTypes"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"/>
                </el-select>
              </el-form-item>
              <el-form-item label="通知配置">
                <el-input v-model="form.desc" type="textarea"/>
              </el-form-item>
            </div>

            <div v-show="activeIndex === 'config-user'">
              <el-form-item label="新建项目用户角色">
                <el-select
                  v-model="selectedNotifyTypes"
                  multiple
                  filterable
                  allow-create
                  default-first-option
                  placeholder="请选择">
                  <el-option
                    v-for="item in selectNotifyTypes"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"/>
                </el-select>
              </el-form-item>
              <el-form-item label="开放注册">
                <el-switch v-model="form.enableRegister"/>
              </el-form-item>
              <el-form-item label="开启 LDAP ">
                <el-switch v-model="form.enableLdap"/>
              </el-form-item>
            </div>
            <div v-show="activeIndex === 'config-project'">
              <el-form-item label="新建项目用户角色">
                <el-select
                  v-model="selectedNotifyTypes"
                  multiple
                  filterable
                  allow-create
                  default-first-option
                  placeholder="请选择">
                  <el-option
                    v-for="item in selectNotifyTypes"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"/>
                </el-select>
              </el-form-item>
            </div>

            <el-form-item>
              <el-button type="primary">保存</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
import apic from '@/api/const'
import { Message } from 'element-ui'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import apiSystem from '@/api/system'

export default {
  name: 'SystemSetting',
  components: { Pagination },
  data() {
    return {
      activeIndex: 'config-deploy',

      form: {
        defaultProjectRoleId: 0,
        enabledRegister: false,
        enabledLdap: false,

        deployRootPath: '',
        customDeployPath: true,
        deployUser: '',
        customDeployUser: true,
        useJenkins: false,
        jenkinsUrl: '',
        jenkinsTemplate: '',
        jenkinsUser: '',
        jenkinsPassword: '',

        pkgLimit: 0,
        pkgStorageType: '',
        pkgStorageConfig: {
          oss: {
            endpoint: '',
            accessKey: '',
            accessSecret: '',
            bucket: ''
          }
        },

        gitCIAuthType: 1,
        gitCIUser: '',
        gitCIPassword: '',
        gitCIToken: '',

        pkgRootPath: '',
        repoRootPath: '',
        buildRootPath: '',
        notifyEnablesTypes: [],
        notifyConfig: {}
      },

      selectedNotifyTypes: [],

      selectNotifyTypes: [{
        value: 'email',
        label: '邮件'
      }, {
        value: 'wechatwork',
        label: '企业微信'
      }],

      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      list: [],
      loadingList: false,
      apic: apic
    }
  },

  created() {
    this.loadSystemConfig()
  },
  methods: {

    setFormValue(sys) {
      this.form.defaultProjectRoleId = sys.default_project_role_id
      this.form.enableRegister = sys.enable_register
      this.form.enableLdap = sys.enable_ldap

      this.form.deployRootPath = sys.deploy_root_path
      this.form.customDeployPath = sys.custom_deploy_path
      this.form.deployUser = sys.deploy_user
      this.form.customDeployUser = sys.custom_deploy_user
      this.form.useJenkins = sys.use_jenkins
      this.form.jenkinsUrl = sys.jenkins_url
      this.form.jenkinsTemplate = sys.jenkins_template
      this.form.jenkinsUser = sys.jenkins_user
      this.form.jenkinsPassword = sys.jenkins_password

      this.form.pkgLimit = sys.pkg_limit
      this.form.pkgStorageType = sys.pkg_storage_type
      this.form.pkgStorageConfig.oss = {
        endpoint: sys.pkg_storage_config.oss.endpoint,
        accessKey: sys.pkg_storage_config.oss.access_key,
        accessSecret: sys.pkg_storage_config.oss.access_secret,
        bucket: sys.pkg_storage_config.oss.bucket
      }

      this.form.gitCIAuthType = sys.git_ci_auth_type
      this.form.gitCIUser = sys.git_ci_user
      this.form.gitCIPassword = sys.git_ci_password
      this.form.gitCIToken = sys.git_ci_token

      this.form.pkgRootPath = sys.pkg_root_path
      this.form.repoRootPath = sys.repo_root_path
      this.form.buildRootPath = sys.build_root_path
      this.form.notifyEnablesTypes = sys.notify_enabled_types.split(',')
      // notifyConfig: {}
    },

    loadSystemConfig() {
      apiSystem.get().then(response => {
        this.system = response.data.system
        this.setFormValue(response.data.system)
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },
    handleSelect(curIndex) {
      this.activeIndex = curIndex
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .dashboard-editor-container {
    margin: 0 auto;
    width: 100%;
    /*min-width: 800px;*/
    padding: 32px;
    background-color: #fff;
    min-height: 550px;
    .chart-wrapper {
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
  .project-manage {
    font-size: 15px;
    margin: 15px 0;
  }
  .perm-item {
    margin: 0 5px 5px 0;
  }

  .status-pass{
    color: #67c23a;
  }
  .status-reject{
    color: red;
  }
  .status-cancel{
    color: #999;
  }
</style>
