<template>
  <div class="app-container">
    <div class="table-out-title"> 选择部署环境 <span class="tail"/></div>

    <el-table
      v-loading="listLoading"
      ref="deployTable"
      :key="tableKey"
      :data="list"
      :row-key="getRowKeys"
      :expand-row-keys="expands"
      border
      fit
      stripe
      style="width: 100%;"
      class="table-primary">
      <el-table-column type="expand">
        <template slot-scope="scope">
          <el-table :data="scope.row.envs">
            <el-table-column property="date" label="部署环境" width="150" align="center">
              <template slot-scope="env">
                {{ env.row.env.name }}
              </template>
            </el-table-column>
            <el-table-column property="name" label="服务器分组" width="200" align="center">
              <template slot-scope="env">
                <span class="link-primary action-link" @click="handleViewCluster(scope.row.app, env.row)">
                  {{ env.row.clusters.length }} 个
                </span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('page.deploy.permitBranches')" property="address" align="center">
              <template slot-scope="env">
                {{ env.row.env.permit_branches.trim() === '*' ? '所有' : env.row.env.permit_branches }}
              </template>
            </el-table-column>
            <el-table-column property="address" label="审核" align="center">
              <template slot-scope="env">
                {{ env.row.env.is_need_audit ? '是' :'否' }}
              </template>
            </el-table-column>
            <el-table-column
              :label="$t('page.codeManagement.action')"
              align="center"
              width="230"
              class-name="small-padding fixed-width">
              <template slot-scope="env">
                <span class="link-primary action-link" @click="handleConfigDeployment(scope.row.app, env.row.env, env.row.clusters)">{{ $t('page.deploy.launchDeploy') }}</span>
              </template>
            </el-table-column>
          </el-table>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.id')" prop="id" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.app.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.app')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.description')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.app.description }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.appType')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ $t('page.deploy.appTypeLabels.' + scope.row.app.app_type) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('page.codeManagement.action')"
        align="center"
        width="230"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handleToggleEnv(scope.row)">{{ $t('page.deploy.env') }} ({{ scope.row.envs.length }})</span>
        </template>
      </el-table-column>
    </el-table>

    <!--View env clusters-->
    <el-dialog :visible.sync="dialogClusterVisible" :title="viewEnvClustersTitle">
      <el-table :data="viewEnvClusters">
        <el-table-column property="name" label="分组名称" />
        <el-table-column property="hosts" label="服务器列表" align="left">
          <template slot-scope="scope">
            <div v-html="scope.row.hosts ? scope.row.hosts.split(',').join('<br>'):'-'"/>
          </template>
        </el-table-column>
      </el-table>
      <div style="height: 35px;"/>
    </el-dialog>

    <el-dialog
      :visible.sync="dialogSelectPkgVisible"
      title="配置部署单"
      top="10vh"
      default-first-option>

      <el-row :gutter="20">
        <el-col :xs="24" :sm="24" :lg="24">
          <div>
            <div >
              <el-tag>应用：{{ deploymentConfig.app.name }}</el-tag>
              <el-tag type="success">环境：{{ deploymentConfig.env.name }}</el-tag>
            </div>
            <p >
              <el-select v-model="deploymentConfig.selectedPkgId" class="filter-item" filterable placeholder="选择版本包" style="width: 100%">
                <el-option
                  v-for="pkg in deploymentConfig.pkgs"
                  :key="pkg.id"
                  :label="pkg.name"
                  :value="pkg.id"/>
              </el-select>
            </p>
            <div style="margin-top: 25px;">
              <div class="table-out-title">  <span class="tail">选择部署分组</span></div>
              <el-table
                ref="deploymentClusterTable"
                :data="deploymentConfig.clusters">
                <el-table-column
                  type="selection"
                  width="55"/>
                <el-table-column property="name" label="分组名称" />
                <el-table-column property="hosts" label="服务器列表" align="left">
                  <template slot-scope="scope">
                    <div v-html="scope.row.hosts ? scope.row.hosts.split(',').join('<br>'):'-'"/>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </el-col>
      </el-row>

      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="handleSubmitDeployment">发起部署</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import apiApp from '@/api/app'
import apiPkg from '@/api/pkg'
import { Message } from 'element-ui'

export default {
  name: 'ComplexTable',
  data() {
    return {
      tableKey: 0,
      list: null,
      listLoading: true,
      dialogPvVisible: false,
      pvData: [],

      getRowKeys(row) {
        return row.app.id
      },

      // rows to expand
      expands: [],

      // envs grid data
      dialogClusterVisible: false,
      viewEnvClusters: [],
      viewEnvClustersTitle: '',

      // select pkg to deploy
      dialogSelectPkgVisible: false,
      deploymentConfig: {
        app: {},
        env: {},
        clusters: [],
        pkgs: {},
        selectedPkgId: ''
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      const project = this.$store.state.project.curProject
      if (!project) {
        Message.error('请先选择项目')
        return
      }
      this.listLoading = true
      apiApp.listAppEnvsByProjectId(project.id).then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.total = response.data.total

        if (this.list.length > 0) {
          this.expands.push(this.list[0].app.id)
        }
      }).catch(error => {
        this.listLoading = false
        console.log(error)
        Message.error(error)
      })
    },

    handleToggleEnv(row) {
      this.$refs.deployTable.toggleRowExpansion(row)
    },

    handleViewCluster(app, row) {
      this.dialogClusterVisible = true
      this.viewEnvClustersTitle = app.name + ':' + row.env.name + ':' + '服务器分组'
      this.viewEnvClusters = row.clusters
    },
    handleConfigDeployment(app, env, clusters) {
      this.dialogSelectPkgVisible = true
      this.deploymentConfig.app = app
      this.deploymentConfig.env = env
      this.deploymentConfig.clusters = clusters
      this.deploymentConfig.pkgs = []
      this.deploymentConfig.selectedPkgId = ''
      apiPkg.listByEnvId(env.id, 20).then(response => {
        this.deploymentConfig.pkgs = response.data.list
        this.$refs.deploymentClusterTable.clearSelection()
        this.$refs.deploymentClusterTable.toggleAllSelection()
      }).catch(error => {
        Message.error(error)
      })
    },
    handleSubmitDeployment() {
      this.dialogSelectPkgVisible = false
      console.log(this.deploymentConfig)
    }
  }
}
</script>
