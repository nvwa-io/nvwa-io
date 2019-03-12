<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" size="mini" icon="el-icon-share" type="primary" @click="handleCreate">
        {{ $t('page.deploy.launchBuild') }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      ref="deployTable"
      :key="tableKey"
      :data="list"
      :row-key="getRowKeys"
      :expand-row-keys="expands"
      border
      stripe
      fit
      style="width: 100%;"
      class="table-primary">
      <el-table-column type="expand">
        <template slot-scope="scope">
          <div style="padding: 15px 0">

            <el-steps :active="getBuildActive(scope.row.build.status)+1" class="step-mini" finish-status="success">
              <el-step v-for="(item, key) in formatBuildSteps(scope.row.build.status)" :key="key" :title="item.title" :status="item.status"/>
            </el-steps>
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.codeManagement.id')" prop="id" align="center" width="65">
        <template slot-scope="scope">
          <!--<span>  {{ scope.$index }}</span>-->
          <span>{{ scope.row.build.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.app')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.branch')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.build.branch }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.commit')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.build.commit_id || scope.row.build.tag || '-' }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.packageName')" min-width="180px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.build.package_name || '-' }}</span>
        </template>
      </el-table-column>
      <!--<el-table-column min-width="150px" align="center">-->
      <el-table-column :label="$t('page.deploy.utime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.build.utime }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.status')" align="center">
        <template slot-scope="scope">
          <p :class="getBuildStatusClass(scope.row.build.status)" >
            <i
              v-if="[
                apic.BUILD_STATUS_CREATED,
                apic.BUILD_STATUS_BUILDING,
                apic.BUILD_STATUS_BUILD_SUCC,
                apic.BUILD_STATUS_PACK_SUCC].indexOf(scope.row.build.status) !== -1"
              class="el-icon-loading"/>
            <span>{{ $t('page.deploy.statusLabels.' + scope.row.build.status) }}</span>
          </p>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('page.codeManagement.action')"
        align="center"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handleToggleEnv(scope.row)">{{ $t('page.deploy.process') }}</span>
          <span class="link-primary action-link" @click="handleBuildLog(scope.row)">{{ $t('page.deploy.log') }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"/>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="450px">
      <el-form
        ref="dataForm"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="margin: auto 20px">

        <el-form-item :label="$t('page.deploy.app')">
          <el-select
            v-model="temp.appId"
            :loading="loadingApps"
            :loading-text="$t('page.deploy.loadingApps')"
            :placeholder="$t('page.deploy.selectApp')"
            filterable
            prop="appId"
            class="filter-item"
            style="width: 100%;"
            @change="changeSelectApp">
            <el-option v-for="item in appsOptions" :key="item.app.id" :label="item.app.name + ' - ' + item.app.description" :value="item.app.id"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('page.deploy.branch')">
          <el-select
            v-model="temp.branch"
            :loading="loadingBranches"
            :loading-text="$t('page.deploy.loadingBranches')"
            :placeholder="$t('page.deploy.selectBranches')"
            filterable
            prop="branch"
            class="filter-item"
            style="width: 100%;">
            <el-option v-for="item in branchesOptions" :key="item" :label="item" :value="item"/>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="createData()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :title="$t('page.deploy.log')" :visible.sync="dialogBuildLogVisible" width="80%" top="5vh">
      <div class="build-log" v-html="buildLog"/>
    </el-dialog>

  </div>
</template>

<script>
import apic from '@/api/const'
import apiBuild from '@/api/build'
import apiApp from '@/api/app'
import { Message } from 'element-ui'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import ShellEditor from '@/components/Editors/shellEditor'

export default {
  name: 'ComplexTable',
  components: { Pagination, ShellEditor },
  data() {
    return {
      tableKey: 0,
      list: null,
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
      branchesOptions: [],
      appsOptions: [],
      showReviewer: false,
      temp: {
        appId: undefined,
        branch: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        create: '发起构建'
      },

      dialogBuildLogVisible: false,
      buildLog: '',
      getRowKeys(row) {
        return row.build.id
      },
      apic: apic,

      expands: [],
      loadingBranches: false,
      loadingApps: false,

      // timer to refresh build status & steps
      timer: undefined
    }
  },
  created() {
    this.getList()
  },
  mounted() {
    if (!this.timer) {
      this.timer = setInterval(() => {
        // check build status, while build status is not end status
        if (!this.list) {
          return
        }

        const length = this.list.length
        for (let i = 0; i < length; i++) {
          // build stats non-end, refresh list
          if ([apic.BUILD_STATUS_BUILD_FAILED,
            apic.BUILD_STATUS_PACK_FAILED,
            apic.BUILD_STATUS_PKG_PUSH_SUCC,
            apic.BUILD_STATUS_PKG_PUSH_FAILED].indexOf(this.list[i].build.status) !== -1) {
            this.getList(false, false)
            break
          }
        }
      }, 5000)
    }
  },
  destroyed() {
    if (this.timer) {
      clearInterval(this.timer)
    }
  },
  methods: {
    getBuildStatusClass(status) {
      switch (status) {
        case apic.BUILD_STATUS_CREATED:
        case apic.BUILD_STATUS_BUILDING:
        case apic.BUILD_STATUS_BUILD_SUCC:
        case apic.BUILD_STATUS_PACK_SUCC:
          return 'status-process'
        case apic.BUILD_STATUS_BUILD_FAILED:
        case apic.BUILD_STATUS_PACK_FAILED:
        case apic.BUILD_STATUS_PKG_PUSH_FAILED:
          return 'status-failed'
        case apic.BUILD_STATUS_PKG_PUSH_SUCC:
          return 'status-success'
        default:
          return ''
      }
    },
    getBuildActive(status) {
      switch (status) {
        case apic.BUILD_STATUS_CREATED:
        case apic.BUILD_STATUS_BUILDING:
        case apic.BUILD_STATUS_BUILD_FAILED:
          return 1
        case apic.BUILD_STATUS_BUILD_SUCC:
        case apic.BUILD_STATUS_PACK_FAILED:
          return 2
        case apic.BUILD_STATUS_PACK_SUCC:
        case apic.BUILD_STATUS_PKG_PUSH_FAILED:
          return 3
        case apic.BUILD_STATUS_PKG_PUSH_SUCC:
          return 4
        default:
          return 0
      }
    },
    formatBuildSteps(status) {
      status = parseInt(status)
      const steps = [
        // ready
        {
          values: [apic.BUILD_STATUS_CREATED],
          success: apic.BUILD_STATUS_CREATED,
          successLabel: '就绪',
          title: '',
          status: ''
        },
        // building, build success, build failed
        {
          values: [apic.BUILD_STATUS_BUILDING, apic.BUILD_STATUS_BUILD_SUCC, apic.BUILD_STATUS_BUILD_FAILED],
          success: apic.BUILD_STATUS_BUILD_SUCC,
          successLabel: '构建通过',
          waitLabel: '构建',
          title: '',
          status: ''
        },
        // packing, pack success, pack failed
        {
          values: [apic.BUILD_STATUS_PACK_SUCC, apic.BUILD_STATUS_PACK_FAILED],
          success: apic.BUILD_STATUS_PACK_SUCC,
          successLabel: '已打包',
          waitLabel: '打包',
          title: '',
          status: ''
        },
        // saving package, save success, save failed
        {
          values: [apic.BUILD_STATUS_PKG_PUSH_SUCC, apic.BUILD_STATUS_PKG_PUSH_FAILED],
          success: apic.BUILD_STATUS_PKG_PUSH_SUCC,
          successLabel: '完成',
          waitLabel: '保存版本包',
          title: '',
          status: ''
        }
      ]

      for (let i = 0; i < steps.length; i++) {
        // in expect status
        let isInExpect = false
        for (let j = 0; j < steps[i].values.length; j++) {
          const v = steps[i].values[j]
          if (v === status) {
            isInExpect = true
            steps[i].title = this.$t('page.deploy.statusLabels.' + status)
            if ([apic.BUILD_STATUS_BUILD_FAILED,
              apic.BUILD_STATUS_PACK_FAILED,
              apic.BUILD_STATUS_PKG_PUSH_FAILED].indexOf(status) !== -1) {
              steps[i].status = 'error'
            } else if (status === apic.BUILD_STATUS_PKG_PUSH_SUCC) {
              steps[i].status = 'success'
            } else {
              steps[i].status = 'process'
            }
          }
        }

        if (isInExpect) {
          continue
        }

        const valuesLen = steps[i].values.length
        if (status > steps[i].values[valuesLen - 1]) {
          steps[i].title = steps[i].successLabel
          steps[i].status = 'success'
        } else {
          steps[i].title = steps[i].waitLabel
          steps[i].status = 'wait'
        }
      }

      // console.log(steps)
      return steps
    },

    // arg1: listLoading, arg2: show message
    getList(...args) {
      const project = this.$store.state.project.curProject
      this.listLoading = args.length > 0 ? args[0] : true
      apiBuild.listByProjectId(project.id).then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.total = response.data.total

        // expand first row
        if (this.expands.length === 0 && this.list.length > 0) {
          this.expands.push(this.list[0].build.id)
        }
      }).catch(error => {
        this.listLoading = false
        if (args.length > 1 && args[1]) {
          Message.error(error)
        }
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },

    loadApps() {
      const project = this.$store.state.project.curProject
      apiApp.listByProjectId(project.id).then(response => {
        this.loadingApps = false
        this.appsOptions = response.data.list

        if (this.appsOptions.length > 0) {
          this.temp.appId = this.appsOptions[0].app.id
          this.loadBranches(this.temp.appId)
        }
      }).catch(error => {
        this.loadingApps = false
        console.log(error)
        Message.error(error)
      })
    },

    // load branches
    loadBranches(appId) {
      this.loadingBranches = true
      apiApp.getBranches(appId).then(response => {
        this.loadingBranches = false
        this.branchesOptions = response.data.list

        if (this.branchesOptions.length > 0) {
          this.temp.branch = this.branchesOptions[0]
        }
      }).catch(error => {
        this.loadingBranches = false
        console.log(error)
        Message.error(error)
      })
    },

    // reload branches while app is change
    changeSelectApp(appId) {
      this.loadBranches(appId)
    },

    resetTemp() {
      this.temp = {
        appId: undefined,
        branch: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
      this.loadApps()
    },
    createData() {
      if (!this.temp.appId || this.temp.appId <= 0) {
        Message.error(this.$t('page.deploy.pleaseSelectApp'))
        return
      }

      if (this.temp.branch === '') {
        Message.error(this.$t('page.deploy.pleaseSelectBranch'))
        return
      }
      apiBuild.create(this.temp.appId, this.temp.branch).then(response => {
        this.dialogFormVisible = false
        Message.success(this.$t('page.deploy.operationSuccess'))
        this.getList()
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },

    handleBuildLog(row) {
      this.dialogBuildLogVisible = true
      this.buildLog = row.build.log.split('\n').join('<br>')
      if (this.buildLog === '') {
        this.buildLog = '-'
      }
    },

    handleToggleEnv(row) {
      this.$refs.deployTable.toggleRowExpansion(row)
    }

  }
}
</script>

<style scoped>
 .status-process {
   color: #4d6889;
 }
 .status-success{
    color: #67c23a;
 }
 .status-failed{
   color: red;
 }

.build-log {
  border-radius: 2px;
  background: #393939;
  padding: 10px 15px;
  line-height: 1.5em;
  color: white;
  /*max-height: 550px;*/
  /*overflow-y: scroll;*/
  min-height: 300px;
}

</style>
