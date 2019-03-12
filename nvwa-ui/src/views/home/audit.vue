<template>
  <div>
    <div class="dashboard-editor-container">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/home' }">{{ $t('page.deploy.home') }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ $t('page.deploy.order') }}</el-breadcrumb-item>
      </el-breadcrumb>
      <br>

      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="wait">{{ $t('page.deploy.waitAudit') }}</el-menu-item>
        <el-menu-item index="mine">{{ $t('page.deploy.myDeploymentOrder') }}</el-menu-item>
        <el-menu-item index="audited">{{ $t('page.deploy.audited') }}</el-menu-item>
      </el-menu>
      <div class="blk"/>

      <el-table
        v-loading="loadingList"
        :data="list"
        border
        fit
        stripe
        style="width: 100%"
        class="table-primary">
        <el-table-column
          :label="$t('page.deploy.id')"
          align="center"
          width="45">
          <template slot-scope="scope">
            {{ scope.row.audit.id }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.projectName')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.project.name }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.appName')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.app.name }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.env')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.env.name }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.packageName')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.deployment.pkg }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.creator')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.user.display_name }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.auditUser')"
          align="center">
          <template slot-scope="scope">
            {{ !scope.row.audit_user.display_name ? '-': scope.row.audit_user.display_name }}
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.deploy.status')"
          align="center">
          <template slot-scope="scope">
            <span :class="getAuditStatusClass(scope.row.audit.status)">
              {{ $t('page.deploy.auditStatusLabels.' + scope.row.audit.status) }}
            </span>
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('page.deploy.ctime')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.project.ctime }}
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('page.action')"
          align="center">
          <template slot-scope="scope">
            <span v-if="scope.row.audit.status === apic.AUDIT_STATUS_WAITING">
              <span v-if="scope.row.has_perm_audit" >
                <span class="link-primary action-link" @click="handlePass(scope.row.audit.id)">{{ $t('page.deploy.pass') }}</span>
                <span class="link-danger action-link" @click="handleReject(scope.row.audit.id)">{{ $t('page.deploy.reject') }}</span>
              </span>

              <span
                v-if="isCurUser(scope.row.user.id)"
                class="link-gray action-link"
                @click="handleCancel(scope.row.audit.id)"
              >{{ $t('page.deploy.cancel') }}</span>
            </span>

            <span v-if="scope.row.audit.status !== apic.AUDIT_STATUS_WAITING">
              -
            </span>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>

import apiAudit from '@/api/audit'
import apic from '@/api/const'
import { Message } from 'element-ui'

export default {
  name: 'AuditManagement',
  components: {},
  data() {
    return {
      activeIndex: 'wait',

      list: [],
      loadingList: false,
      apic: apic
    }
  },
  created() {
    this.getList()
  },
  methods: {
    isCurUser(uid) {
      return this.$store.state.user.curUser.id === uid
    },
    getList() {
      this.getListByMenu()
    },
    handleSelect(index, indexPath) {
      this.activeIndex = index
      this.getListByMenu()
    },

    getListByMenu() {
      let req
      switch (this.activeIndex) {
        case 'wait':
          req = apiAudit.listWait()
          break
        case 'mine':
          req = apiAudit.listMine()
          break
        case 'audited':
          req = apiAudit.listAudited()
          break
        default:
          Message.error('Invalid index.')
          return
      }

      this.loadingList = true
      req.then(response => {
        this.loadingList = false
        this.list = response.data.list
      }).catch(error => {
        this.loadingList = false
        console.log(error)
        Message.error(error)
      })
    },
    handlePass(auditId) {
      this.$confirm(this.$t('page.deploy.tipPassAudit'), this.$t('page.deploy.titlePassAudit'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiAudit.pass(auditId).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
          this.$store.dispatch('GetWaitAuditNum')
        }).catch(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    },
    handleReject(auditId) {
      this.$confirm(this.$t('page.deploy.tipRejectAudit'), this.$t('page.deploy.titleRejectAudit'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiAudit.reject(auditId).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
          this.$store.dispatch('GetWaitAuditNum')
        }).catch(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    },
    handleCancel(auditId) {
      this.$confirm(this.$t('page.deploy.tipCancelAudit'), this.$t('page.deploy.titleCancelAudit'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiAudit.cancel(auditId).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
          this.$store.dispatch('GetWaitAuditNum')
        }).catch(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    },

    getAuditStatusClass(status) {
      switch (status) {
        case apic.AUDIT_STATUS_WAITING:
          return ''
        case apic.AUDIT_STATUS_PASS:
          return 'status-pass'
        case apic.AUDIT_STATUS_REJECT:
          return 'status-reject'
        case apic.AUDIT_STATUS_CANCELED:
          return 'status-cancel'
        default:
          return ''
      }
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
