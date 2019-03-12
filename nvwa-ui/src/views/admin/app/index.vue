<template>
  <div class="app-container">
    <div class="table-out-title"> {{ $t('page.admin.deploy.manageApp') }}<span class="tail"/></div>

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

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"/>

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
  </div>

</template>

<script>
import { updateArticle } from '@/api/article'
import apiSystem from '@/api/system'
import apiApp from '@/api/admin/app'
import { Message } from 'element-ui'
import waves from '@/directive/waves' // Waves directive
import ShellEditor from '@/components/Editors/shellEditor'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ComplexTable',
  components: { ShellEditor, Pagination },
  directives: { waves },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },

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

      // @TODO 待删除
      pvData: [],
      rules: {
        name: [{ required: true, message: '应用名称不能为空', trigger: 'blur' }],
        deploy_user: [{ required: true, message: '部署用户不能为空', trigger: 'blur' }],
        deployPath: [{ required: true, message: '部署路径不能为空', trigger: 'blur' }],
        repo_url: [{ required: true, message: 'Git 仓库不能为空', trigger: 'blur' }]
      },
      downloadLoading: false,

      gridData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }],

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
      apiApp.list().then(response => {
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
          console.log(this.formAppData)
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
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          updateArticle(tempData).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
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
