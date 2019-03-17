<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleCreate">
        {{ $t('page.deploy.addApp') }}
      </el-button>
      <el-button class="filter-item btn-blank" size="mini" icon="el-icon-question" @click="handleViewHelp">
        {{ $t('page.viewHelp') }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      stripe
      highlight-current-row
      style="width: 100%;"
      class="table-primary">
      <el-table-column :label="$t('page.deploy.id')" prop="id" align="center" width="45">
        <template slot-scope="scope">
          <span>{{ scope.row.app.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.app')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.description')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.description }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.creator')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.user.username }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.utime')" min-width="90" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.utime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.buildDeploy')" min-width="90" align="center">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handleBuildCmd(scope.row.app)">{{ $t('page.deploy.buildCmd') }}</span>
          <span class="link-primary action-link" @click="handleDeployCmd(scope.row.app)">{{ $t('page.deploy.deployCmd') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.env')" min-width="90" align="center" >
        <template slot-scope="scope">
          <!--<span class="link-primary action-link" @click="handleUpdate(scope.row)">{{ $t('page.deploy.groupHost') }}</span>-->
          <router-link :to="{path:'/app-list/env', query:{app_id: scope.row.app.id}}" class="link-primary action-link">{{ $t('page.deploy.env') }}</router-link>
          <!--<router-link class="link-primary action-link" to="/app-list/env">{{ $t('page.deploy.appStatus') }}</router-link>-->
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('page.codeManagement.action')"
        align="center"
        width="120"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handleUpdate(scope.row)">{{ $t('page.opTable.edit') }}</span>
          <span
            class="link-danger action-link"
            @click="handleUpdate(scope.row)">{{ $t('page.codeManagement.delete') }}</span>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="850px" top="50px">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="formAppData"
        label-position="left"
        label-width="90px"
        style="padding: 0 20px">
        <el-row :gutter="30">
          <el-col :xs="12" :sm="12" :lg="12">
            <el-form-item :label="$t('page.deploy.appName')" prop="name">
              <el-input v-model="formAppData.name" :disabled="formAppData.id > 0" placeholder="e.g: nvwa-server" />
            </el-form-item>
            <el-form-item :label="$t('page.deploy.appDescription')" prop="description">
              <el-input
                :autosize="{ minRows: 2, maxRows: 4}"
                v-model="formAppData.description"
                type="textarea"
                placeholder="e.g: 女娲服务端"/>
            </el-form-item>
            <el-form-item :label="$t('page.deploy.deployUser')" prop="deploy_user">
              <el-input v-model="formAppData.deploy_user" :value="system.deploy_user" :disabled="!system.custom_deploy_user" placeholder="e.g: nvwa"/>
            </el-form-item>
            <el-form-item :label="$t('page.deploy.deployPath')" prop="deployPath">
              <el-input v-model="formAppData.deploy_path" :disabled="!system.custom_deploy_path" placeholder="e.g: nvwa-server"/>
            </el-form-item>
          </el-col>
          <el-col :xs="12" :sm="12" :lg="12">
            <el-form-item :label="$t('page.deploy.gitRepo')" prop="repo_url" >
              <el-input v-model="formAppData.repo_url" placeholder="e.g: http://github.com/nvwa-io/nvwa-io"/>
            </el-form-item>

            <div v-if="!httpConfigVisible" class="warning">
              <p>备注：支持 <el-tag type="success" size="mini">SSH</el-tag> 和 <el-tag size="mini">HTTP(S)</el-tag></p>
              <p><el-tag type="success" size="mini">SSH</el-tag> {{ $t('page.deploy.gitSshTip') }}</p>
              <p><el-tag size="mini">HTTP(S)</el-tag> {{ $t('page.deploy.gitHttpTip') }}</p>
            </div>
            <div>
              <div v-if="httpConfigVisible">
                <p><el-tag size="mini">HTTP(S)</el-tag> {{ $t('page.deploy.gitHttpTip') }}</p>
                <el-radio-group v-model="httpConfig" size="mini">
                  <el-radio label="system" border>系统默认</el-radio>
                  <el-radio label="custom" border>自定义</el-radio>
                </el-radio-group>
                <div v-if="httpConfig === 'custom'" class="blk">
                  <el-form-item :label="$t('page.deploy.gitUsername')" prop="gitUsername">
                    <el-input v-model="formAppData.repo_username" placeholder="e.g: nvwa-server"/>
                  </el-form-item>
                  <el-form-item :label="$t('page.deploy.gitPassword')" prop="gitPassword">
                    <el-input v-model="formAppData.repo_password" placeholder="e.g: nvwa-server"/>
                  </el-form-item>
                </div>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingAppForm" type="primary" @click="dialogStatus==='create'?createData():updateData()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogHelp" title="帮助" width="70%">
      <el-steps :active="5" align-center style="margin-bottom: 30px;">
        <el-step title="步骤1" description="新建应用"/>
        <el-step title="步骤2" description="配置部署环境"/>
        <el-step title="步骤3" description="配置构建命令"/>
        <el-step title="步骤4" description="配置部署命令"/>
        <el-step title="步骤5" description="完成"/>
      </el-steps>
      <div class="blk">
        <p class="text-center text-hint">新应用配置流程</p>
      </div>
    </el-dialog>

    <!--:close-on-click-modal="false"-->
    <el-dialog :visible.sync="dialogDeployCmd" title="部署命令" width="85%" top="5vh">
      <el-steps :active="-1" align-center class="step-mini" style="margin-bottom: 30px;">
        <el-step title="步骤1" description="检测基础环境"/>
        <el-step title="步骤2" description="服务器下载和解压版本包"/>
        <el-step title="步骤4" status="process" description="执行部署前命令"/>
        <el-step title="步骤5" description="软链版本至部署目录"/>
        <el-step title="步骤6" status="process" description="执行部署后命令"/>
        <el-step title="步骤7" status="process" description="应用检测"/>
        <el-step title="步骤8" status="process" description="执行上线命令"/>
        <el-step title="步骤9" description="完成"/>
      </el-steps>
      <div class="blk" style="margin-bottom: 30px;">
        <p class="text-center text-hint">应用部署流程</p>
      </div>

      <el-tabs v-model="activeCmdTabName">
        <el-tab-pane label="部署前命令（可选）" name="cmd_before_deploy"/>
        <el-tab-pane label="部署后命令（可选）" name="cmd_after_deploy"/>
        <el-tab-pane label="应用检测（可选）" name="cmd_health_check"/>
        <el-tab-pane label="上线命令（可选）" name="cmd_online"/>
      </el-tabs>

      <div>
        <el-row :gutter="20">
          <el-col :xs="18" :sm="18" :lg="18">
            <div class="editor-container">
              <shell-editor ref="shellEditor" v-model="tmpEditorCmd"/>
            </div>
          </el-col>
          <el-col :xs="6" :sm="6" :lg="6">
            <div class="table-out-title">超时时间</div>
            <el-input v-model="formCmds.cmd_timeout" placeholder="e.g: 3600">
              <template slot="append">秒</template>
            </el-input>
          </el-col>
        </el-row>
      </div>

      <div style="color: #999;">
        <p>备注：所有命令（shell 命令）都是在版本的根目录下执行，另外部署系统提供一批变量供使用——<a class="link-type" href="">查看变量</a></p>
      </div>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogDeployCmd= false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingDeployCmd" type="primary" @click="updateCmdSubmit">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <!--:close-on-click-modal="false" -->
    <el-dialog :visible.sync="dialogBuildCmd" title="构建命令" width="85%" top="5vh">
      <el-steps :active="-1" align-center class="step-mini" style="margin-bottom: 30px;">
        <el-step title="步骤1" description="拉取/更新代码"/>
        <el-step title="步骤2" description="初始化构建工作空间"/>
        <el-step title="步骤4" status="process" description="执行构建命令"/>
        <el-step title="步骤5" status="process" description="打版本包"/>
        <el-step title="步骤6" description="保存版本包"/>
        <el-step title="步骤7" description="完成"/>
      </el-steps>
      <div class="blk" style="margin-bottom: 30px;">
        <p class="text-center text-hint">应用构建和打包流程</p>
      </div>

      <el-tabs v-model="activeCmdTabName">
        <el-tab-pane label="构建命令（可选）" name="cmd_build"/>
      </el-tabs>

      <div>
        <el-row :gutter="20">
          <!--<el-col :xs="15" :sm="15" :lg="15">-->
          <el-col :xs="24" :sm="24" :lg="24">
            <div class="editor-container">
              <shell-editor ref="shellEditor" v-model="formBuilds.cmd_build"/>
            </div>
          </el-col>
        </el-row>

        <el-row :gutter="30" style="margin-top: 20px">
          <el-col :xs="12" :sm="12" :lg="12">
            <div class="table-out-title">配置打包的文件（或目录）</div>
            <shell-editor ref="shellEditor" v-model="formBuilds.files" />
          </el-col>
          <el-col :xs="12" :sm="12" :lg="12">
            <div class="table-out-title">配置排除打包的文件（或目录）</div>
            <shell-editor ref="shellEditor" v-model="formBuilds.excludes"/>
          </el-col>
        </el-row>
        <div style="color: #999;">
          <p>备注：打包命名的工作目录跟构建命令执行完的最终工作目录一致。</p>
        </div>
      </div>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogBuildCmd= false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingBuildCmd" type="primary" @click="updateBuildSubmit">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

  </div>

</template>

<script>
import apiSystem from '@/api/system'
import apiApp from '@/api/app'
import { Message } from 'element-ui'
import waves from '@/directive/waves' // Waves directive
import ShellEditor from '@/components/Editors/shellEditor'

export default {
  name: 'ComplexTable',
  components: { ShellEditor },
  directives: { waves },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      dialogHelp: false,

      formAppData: {
        id: 0,
        name: '',
        description: '',
        deploy_user: '',
        deploy_path: '',
        repo_url: '',
        repo_username: '',
        repo_password: ''
      },
      // system config
      system: {
        deploy_root_path: '',
        deploy_user: '',
        custom_deploy_user: false,
        custom_deploy_path: false
      },
      httpConfigVisible: false,
      httpConfig: 'system',

      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '修改应用',
        create: '新建应用'
      },

      loadingAppForm: false,

      rules: {
        name: [{ required: true, message: '应用名称不能为空', trigger: 'blur' }],
        deploy_user: [{ required: true, message: '部署用户不能为空', trigger: 'blur' }],
        deploy_path: [{ required: true, message: '部署路径不能为空', trigger: 'blur' }],
        repo_url: [{ required: true, message: 'Git 仓库不能为空', trigger: 'blur' }]
      },
      downloadLoading: false,

      dialogDeployCmd: false,
      tmpEditorCmd: '',
      activeCmdTabName: 'cmd_before_deploy',
      formCmds: {
        id: 0,
        cmd_before_deploy: '',
        cmd_after_deploy: '',
        cmd_health_check: '',
        cmd_online: '',
        cmd_timeout: 3600
      },
      loadingDeployCmd: false,

      // build
      dialogBuildCmd: false,
      loadingBuildCmd: false,
      formBuilds: {
        id: 0,
        cmd_build: '',
        files: '',
        excludes: ''
      }
    }
  },
  computed: {
    formAppDataName() {
      return this.formAppData.name
    },
    formAppDataRepoUrl() {
      return this.formAppData.repo_url
    },
    activeCmdTabNameTmp() {
      return this.activeCmdTabName
    }

  },
  watch: {
    formAppDataName(v, oldv) {
      if (this.dialogStatus !== 'create') { // only watch create form
        return
      }

      if (!this.system.custom_deploy_path) { // not allow custom path
        this.formAppData.deploy_path = this.system.deploy_root_path.trimRight().replace(/\/$/, '') + '/' + v
      } else { // allow custom path
        if (!this.formAppData.deploy_path) {
          return
        }

        // if (this.formAppData.deploy_path !== this.system.deploy_root_path) {
        //   return
        // }
        const tmpArr = this.formAppData.deploy_path.split('/')
        if (oldv) {
          tmpArr.pop()
        }
        this.formAppData.deploy_path = tmpArr.join('/').trimRight().replace(/\/$/, '') + '/' + this.formAppData.name
      }
    },
    formAppDataRepoUrl(v, oldv) {
      if (!v) {
        this.httpConfigVisible = false
        return
      }
      if (v.startsWith('http')) {
        this.httpConfigVisible = true
      } else if (v.startsWith('ssh')) {
        this.httpConfigVisible = false
      } else {
        this.httpConfigVisible = false
      }
    },
    activeCmdTabNameTmp(v, oldv) {
      console.log(v, oldv)
      this.formCmds[oldv] = this.tmpEditorCmd
      this.tmpEditorCmd = this.formCmds[v]
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      const project = this.$store.state.project.curProject
      apiApp.listByProjectId(project.id).then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.total = response.data.total
      }).catch(error => {
        this.listLoading = false
        Message.error(error)
      })
    },

    handleModifyStatus(row, status) {
      this.$message({
        message: '操作成功',
        type: 'success'
      })
      row.status = status
    },
    resetFormAppData() {
      this.formAppData = {
        id: 0,
        name: '',
        description: '',
        deploy_user: '',
        deploy_path: '',
        repo_url: '',
        repo_username: '',
        repo_password: ''
      }
    },
    handleCreate() {
      this.resetFormAppData()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true

      apiSystem.get().then(response => {
        this.system = response.data.system
        this.formAppData.deploy_user = this.system.deploy_user
        this.formAppData.deploy_path = this.system.deploy_root_path.trimRight('/')
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.loadingAppForm = true
          apiApp.create({ project_id: this.$store.state.project.curProject.id,
            name: this.formAppData.name,
            description: this.formAppData.description,
            deploy_user: this.formAppData.deploy_user,
            deploy_path: this.formAppData.deploy_path,
            repo_url: this.formAppData.repo_url,
            repo_username: this.formAppData.repo_username,
            repo_password: this.formAppData.repo_password }).then(response => {
            this.dialogFormVisible = false
            this.loadingAppForm = false
            console.log(response)
            Message.success('成功新建应用')
            this.fetchList()
          }).catch(error => {
            this.loadingAppForm = false
            console.log(error)
            Message.error(error)
          })
        }
      })
    },

    handleViewHelp() {
      this.dialogHelp = true
    },
    handleBuildCmd(app) {
      this.dialogBuildCmd = true

      // init cmd values
      this.formBuilds = {
        id: app.id,
        cmd_build: app.cmd_build,
        files: app.files,
        excludes: app.excludes
      }
    },

    handleDeployCmd(app) {
      this.dialogDeployCmd = true
      // reset command editor and tab
      this.activeCmdTabName = 'cmd_before_deploy'
      this.tmpEditorCmd = app.cmd_before_deploy

      // init cmd values
      this.formCmds = {
        id: app.id,
        cmd_before_deploy: app.cmd_before_deploy,
        cmd_after_deploy: app.cmd_after_deploy,
        cmd_health_check: app.cmd_health_check,
        cmd_online: app.cmd_online,
        cmd_timeout: app.cmd_timeout
      }
    },
    // update deploy commands
    updateCmdSubmit() {
      // update current tab's editor' content
      this.formCmds[this.activeCmdTabName] = this.tmpEditorCmd

      this.loadingDeployCmd = true
      apiApp.updateCmds(this.formCmds.id, this.formCmds).then(response => {
        Message.success('操作成功')
        this.dialogDeployCmd = false
        this.loadingDeployCmd = false
        this.getList()
      }).catch(error => {
        this.loadingDeployCmd = false
        console.log(error)
        Message.error(error)
      })
    },

    // update build config
    updateBuildSubmit() {
      console.log(this.formBuilds)
      this.loadingBuildCmd = true
      apiApp.updateCmds(this.formBuilds.id, this.formBuilds).then(response => {
        Message.success('操作成功')
        this.dialogBuildCmd = false
        this.loadingBuildCmd = false
        this.getList()
      }).catch(error => {
        this.loadingBuildCmd = false
        console.log(error)
        Message.error(error)
      })
    },

    handleUpdate(row) {
      this.formAppData = Object.assign({}, row.app) // copy obj
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          console.log(this.$refs['dataForm'])
        }
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .CodeMirror-gutter-wrapper {
    left: 35px!important;
  }
  </style>
