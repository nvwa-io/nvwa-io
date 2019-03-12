<template>
  <div class="app-container">

    <div class="filter-container">
      <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleCreate">
        {{ $t('page.deploy.addEnv') }}
      </el-button>
      <el-button class="filter-item btn-blank" size="mini" icon="el-icon-question" @click="handleViewHelp">
        {{ $t('page.viewHelp') }}
      </el-button>
    </div>

    <div class="table-out-title"> 环境列表 <span class="tail">· {{ app.name }}</span></div>

    <el-table
      v-loading="listLoading"
      ref="clusterTable"
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
          <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleClusterCreate(scope.row.env)">
            {{ $t('page.deploy.addCluster') }}
          </el-button>
          <div class="blk"/>
          <el-table :data="scope.row.clusters">
            <el-table-column :label="$t('page.deploy.clusterName')" property="name" width="110" align="center">
              <template slot-scope="cluster">
                {{ cluster.row.name }}
              </template>
            </el-table-column>
            <el-table-column width="50" align="center"/>
            <el-table-column :label="$t('page.deploy.clusterHosts')" property="date" width="180" align="left">
              <template slot-scope="cluster">
                <!--<span class="link-primary">{{ cluster.row.hosts === '' ? 0:cluster.row.hosts.split(',').length }}</span>-->
                <span v-html=" cluster.row.hosts === '' ? '-':cluster.row.hosts.split(',').join('<br>')"/>
              </template>
            </el-table-column>
            <el-table-column :label="$t('page.deploy.ctime')" property="name" align="center">
              <template slot-scope="cluster">
                <span>{{ cluster.row.ctime }}</span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('page.deploy.detectionCond')" align="center">
              <template slot-scope="cluster">
                <span class="link-primary action-link" @click="handleUpdate(cluster.row)">
                  <i class="el-icon-rank"/> {{ $t('page.deploy.detection') }}
                </span>
              </template>
            </el-table-column>
            <el-table-column
              :label="$t('page.codeManagement.action')"
              align="center"
              width="160"
              class-name="small-padding fixed-width">
              <template slot-scope="cluster">
                <span class="link-primary action-link" @click="handleClusterUpdate(cluster.row)">{{ $t('page.opTable.edit') }}</span>
                <span class="link-danger action-link" @click="handleClusterUpdate(cluster.row)">{{ $t('page.codeManagement.delete') }}</span>
              </template>
            </el-table-column>
          </el-table>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.id')" prop="id" align="center" width="45">
        <template slot-scope="scope">
          <span>{{ scope.row.env.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.env')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.env.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.cluster')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handlerToggleClusters(scope.row)">{{ scope.row.clusters.length }} 个</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.permitBranches')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.env.permit_branches === '*' ? '所有' : scope.row.env.permit_branches }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.audit')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.env.is_need_audit ? '是':'否' }}</span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('page.codeManagement.action')"
        align="center"
        width="160"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handlerToggleClusters(scope.row)">{{ $t('page.deploy.configCluster') }}</span>
          <span class="link-primary action-link" @click="handleUpdate(scope.row)">{{ $t('page.opTable.edit') }}</span>
          <span
            class="link-danger action-link"
            @click="handleUpdate(scope.row)">{{ $t('page.codeManagement.delete') }}</span>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" top="5vh">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="formEnv"
        label-position="left"
        label-width="90px"
        style="padding: 0 20px;">
        <el-form-item :label="$t('page.deploy.envName')" prop="name">
          <el-input v-model="formEnv.name" placeholder="e.g: 线上环境"/>
        </el-form-item>
        <el-form-item :label="$t('page.deploy.permitBranches')" prop="permit_branches">
          <el-input v-model="formEnv.permit_branches" placeholder="e.g: * 表示全部，多个分支用逗号分隔"/>
        </el-form-item>
        <el-form-item :label="$t('page.deploy.enableAudit')">
          <el-switch v-model="formEnv.is_need_audit"/>
        </el-form-item>
        <div class="blk">
          <p style="font-weight: bold;">环境差异化命令（可选）</p>
          <div style="color: #999;font-size: 14px;line-height: 1.6em;">
            <p>备注：[环境差异化命令] 用于不同环境部署时，自动拼接在部署命令前进行执行（可用于，如：设置系统环境变量等）—— <a class="link-type" href="">查看可用变量</a></p>
          </div>
          <shell-editor ref="shellEditor" v-model="formEnv.cmd_env" height="100px"/>
        </div>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingFormEnv" type="primary" @click="dialogStatus==='create'?createData():updateData()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :title="textClusterMap[dialogClusterStatus]" :visible.sync="dialogClusterVisible" top="5vh">
      <el-form
        ref="clusterForm"
        :rules="clusterRules"
        :model="formCluster"
        label-position="left"
        label-width="100px"
        style="padding: 0 20px;">
        <el-form-item :label="$t('page.deploy.clusterName')" prop="name">
          <el-input v-model="formCluster.name" placeholder="e.g: 集群 01"/>
        </el-form-item>
        <el-form-item :label="$t('page.deploy.clusterHosts')" prop="hosts">
          <div class="form-editor">
            <shell-editor ref="shellEditor" v-model="formCluster.hosts" />
            <div style="color: #999;">
              <p>备注：一个服务器地址一行。</p>
            </div>
          </div>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogClusterVisible= false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingFormCluster" type="primary" @click="dialogClusterStatus==='create'?createClusterData():updateClusterData()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogHelp" title="帮助">
      <el-steps :active="4" align-center style="margin-bottom: 30px;">
        <el-step title="步骤1" description="新建环境"/>
        <el-step title="步骤2" description="服务器分组"/>
        <el-step title="步骤3" description="服务器条件检测"/>
        <el-step title="步骤4" description="完成"/>
      </el-steps>
      <div class="blk">
        <p class="text-center text-hint">新环境配置流程</p>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import apiApp from '@/api/app'
import apiEnv from '@/api/env'
import apiCluster from '@/api/cluster'
import { Message } from 'element-ui'
import ShellEditor from '@/components/Editors/shellEditor'

export default {
  name: 'AppCluster',
  components: { ShellEditor },
  data() {
    return {
      app: {},

      tableKey: 0,
      list: null,
      listLoading: true,
      formEnv: {
        id: undefined,
        app_id: parseInt(this.$route.query.app_id),
        name: '',
        permit_branches: '*',
        is_need_audit: false,
        cmd_env: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '修改环境',
        create: '新建环境'
      },
      rules: {
        name: [{ required: true, message: '环境名称不能为空', trigger: 'change' }],
        permit_branches: [{ required: true, message: '可选分支不能为空', trigger: 'change' }]
      },

      loadingFormEnv: false,

      // cluster
      dialogClusterVisible: false,
      loadingFormCluster: false,
      textClusterMap: {
        update: '修改分组',
        create: '新建分组'
      },
      dialogClusterStatus: '',
      formCluster: {
        id: undefined,
        app_id: parseInt(this.$route.query.app_id),
        env_id: 0,
        name: '',
        hosts: ''
      },
      clusterRules: {
        name: [{ required: true, message: '集群名称不能为空', trigger: 'change' }],
        hosts: [{ required: true, message: '服务器列表不能为空', trigger: 'change' }]
      },
      dialogHelp: false,
      getRowKeys(row) {
        return row.env.id
      },
      expands: []
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      apiApp.listAppEnvsByAppId(this.$route.query.app_id).then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.app = response.data.app
        if (this.expands.length === 0 && this.list.length > 0) {
          this.expands.push(this.list[0].env.id)
        }
      }).catch(error => {
        this.listLoading = false
        console.log(error)
        Message.error(error)
      })
    },
    resetFormEnv() {
      this.formEnv = {
        id: undefined,
        app_id: parseInt(this.$route.query.app_id),
        name: '',
        permit_branches: '*',
        is_need_audit: false,
        cmd_env: ''
      }
    },
    handleCreate() {
      this.resetFormEnv()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.loadingFormEnv = true
          apiEnv.create(this.formEnv).then(response => {
            this.loadingFormEnv = false
            this.dialogFormVisible = false
            Message.success('新建环境成功')
            this.getList()
          }).catch(error => {
            this.loadingFormEnv = false
            console.log(error)
            Message.error(error)
          })
        }
      })
    },

    handleViewHelp() {
      this.dialogHelp = true
    },

    handlerToggleClusters(row) {
      this.$refs.clusterTable.toggleRowExpansion(row)
    },

    handleUpdate(row) {
      this.formEnv = Object.assign({}, row.env)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.loadingFormEnv = true
          apiEnv.update(this.formEnv.id, {
            name: this.formEnv.name,
            permit_branches: this.formEnv.permit_branches,
            is_need_audit: this.formEnv.is_need_audit,
            cmd_env: this.formEnv.cmd_env
          }).then(response => {
            this.loadingFormEnv = false
            this.dialogFormVisible = false
            Message.success('修改环境成功')
            this.getList()
          }).catch(error => {
            this.loadingFormEnv = false
            console.log(error)
            Message.error(error)
          })
        }
      })
    },

    resetClusterFormData() {
      this.formCluster = {
        id: undefined,
        app_id: parseInt(this.$route.query.app_id),
        env_id: 0,
        name: '',
        hosts: ''
      }
    },
    handleClusterCreate(env) {
      this.resetClusterFormData()
      this.dialogClusterVisible = true
      this.formCluster.env_id = env.id
      this.dialogClusterStatus = 'create'
      this.$nextTick(() => {
        this.$refs['clusterForm'].clearValidate()
      })
    },
    handleClusterUpdate(row) {
      this.dialogClusterVisible = true
      this.dialogClusterStatus = 'update'
      this.formCluster = Object.assign({}, row)
      this.formCluster.hosts = this.formCluster.hosts.split(',').join('\n')
      this.$nextTick(() => {
        this.$refs['clusterForm'].clearValidate()
      })
    },
    createClusterData() {
      this.$refs['clusterForm'].validate((valid) => {
        if (valid) {
          this.loadingFormCluster = true
          apiCluster.create({
            app_id: this.formCluster.app_id,
            env_id: this.formCluster.env_id,
            name: this.formCluster.name,
            hosts: this.formCluster.hosts.split('\n').join(',')
          }).then(response => {
            this.loadingFormCluster = false
            this.dialogClusterVisible = false
            Message.success('新建分组成功')
            this.getList()
          }).catch(error => {
            this.loadingFormCluster = false
            console.log(error)
            Message.error(error)
          })
        }
      })
    },
    updateClusterData() {
      this.$refs['clusterForm'].validate((valid) => {
        if (valid) {
          this.loadingFormCluster = true
          apiCluster.update(this.formCluster.id, {
            app_id: this.formCluster.app_id,
            env_id: this.formCluster.env_id,
            name: this.formCluster.name,
            hosts: this.formCluster.hosts.split('\n').join(',')
          }).then(response => {
            this.loadingFormCluster = false
            this.dialogClusterVisible = false
            Message.success('修改分组成功')
            this.getList()
          }).catch(error => {
            this.loadingFormCluster = false
            console.log(error)
            Message.error(error)
          })
        }
      })
    }

    // handleDelete(row) {
    //   this.$notify({
    //     title: '成功',
    //     message: '删除成功',
    //     type: 'success',
    //     duration: 2000
    //   })
    //   const index = this.list.indexOf(row)
    //   this.list.splice(index, 1)
    // }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .ungroup {
    padding: 8px 10px;
    margin-bottom: 10px;
    background: #f7f7f7;
    border-radius: 2px;
    cursor: pointer;
    font-size: 13px;
  }

  .ungroup:hover {
    background: #f0f0f0;
  }

  .host-num {
    display: inline-block;
    padding: 3px 6px 2px;
    border-radius: 13px;
    background: #e5e5e5;
    font-size: 12px;
    margin-left: 5px;
  }

  .group-active {
    color: #105cbc !important;
  }

  .env-box {
    border-radius: 2px;
    margin-bottom: 10px;
    .env-title {
      padding: 8px 10px;
      color: #5a5c60;
      background: #dfedff;
      cursor: pointer;
      font-size: 13px;

      .host-num {
        background: #c4dcf3;
      }
    }
    .env-title:hover {
      background: #d4e6ff;

    }

    .env-group {
      .env-group-item {
        padding: 8px 10px 8px 20px;
        color: #5a5c60;
        border-top: 1px solid #e7f0fa;
        background: #eff6ff;
        cursor: pointer;
        font-size: 13px;

        .host-num {
          background: #dce5ef;
        }
      }
      .env-group-item:hover {
        background: #e8f0ff;
      }
    }
  }

  .op-btn {
    margin: 1px 0 0 10px;
  }
  .form-editor *{
    line-height: 1.1em!important;
  }
</style>
