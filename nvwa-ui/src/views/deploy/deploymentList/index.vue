<template>
  <div class="app-container">
    <!--<div>-->
    <!--<el-select v-model="value" placeholder="所有应用">-->
    <!--<el-option-->
    <!--v-for="item in options"-->
    <!--:key="item.value"-->
    <!--:label="item.label"-->
    <!--:value="item.value"/>-->
    <!--</el-select>-->
    <!--</div>-->

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
      class="table-primary blk">
      <el-table-column type="expand">
        <template slot-scope="scope">
          <!--ref="clusterTable"-->
          <el-table
            :ref="jobPrefix + scope.row.deployment.id"
            :data="scope.row.jobs"
            :row-key="getJobRowKeys"
            :expand-row-keys="jobExpands"
            @expand-change="jobDetailStepsExpandChange">
            <el-table-column type="expand">
              <template slot-scope="job">
                <div style="padding: 15px 0">
                  <el-steps :active="jobDetailStepsActive[job.row.job.id]" class="step-mini" finish-status="success" align-center>
                    <el-step v-for="(item, key) in jobDetailSteps[job.row.job.id]" :key="key" :title="item.label" :status="item.status"/>
                  </el-steps>
                </div>
              </template>
            </el-table-column>

            <el-table-column property="name" label="服务器分组" width="110" align="center">
              <template slot-scope="job">
                <span class="link-primary action-link" @click="handleViewClusterHosts(job.row)">
                  {{ job.row.cluster.name }} ({{ job.row.job.deploy_hosts === '' ? 0 : job.row.job.deploy_hosts.split(',').length }})
                </span>
              </template>
            </el-table-column>
            <el-table-column property="date" label="进度" align="center">
              <template slot-scope="job">
                <el-steps :active="getJobActive(job.row.job.status) + 1" class="step-mini in-column-fix" align-center>
                  <el-step v-for="(item, key) in formatJobShortSteps(job.row.job.status)" :key="key" :title="item.title" :status="item.status"/>
                </el-steps>
              </template>
            </el-table-column>
            <el-table-column property="name" label="操作时间" width="145" align="center">
              <template slot-scope="job">
                <span>{{ job.row.job.utime }}</span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('page.deploy.status')" property="name" width="90" align="center">
              <template slot-scope="job">
                <p :class="getJobStatusClass(job.row.job.status)">
                  <i v-if="[apic.JOB_STATUS_READY, apic.JOB_STATUS_DEALING].indexOf(job.row.job.status) !== -1" class="el-icon-loading"/>
                  {{ $t('page.deploy.jobStatusLabels.' + job.row.job.status) }}
                </p>
              </template>
            </el-table-column>
            <el-table-column
              :label="$t('page.codeManagement.action')"
              align="center"
              width="160"
              class-name="small-padding fixed-width">
              <template slot-scope="job">
                <span class="link-primary action-link" @click="handleToggleJobDetailSteps(job.row)">进度</span>
                <span
                  v-if="[apic.JOB_STATUS_CREATED].indexOf(job.row.job.status) === -1"
                  class="link-primary action-link"
                  @click="handleViewLog(job.row.job)">
                  {{ $t('page.deploy.log') }}
                </span>
                <span
                  v-if="job.row.job.status === apic.JOB_STATUS_CREATED"
                  class="link-primary action-link"
                  @click="handleStartDeploy(job.row.job)">
                  {{ $t('page.deploy.startDeploy') }}
                </span>
              </template>
            </el-table-column>
          </el-table>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.id')" prop="id" width="65" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.deployment.id }}</span>
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
      <el-table-column :label="$t('page.deploy.status')" align="center">
        <template slot-scope="scope">
          <span :class="getDeploymentStatusClass(scope.row.deployment.status)">{{ $t('page.deploy.deploymentStatusLabels.'+scope.row.deployment.status) }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.utime')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.deployment.utime }}</span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('page.codeManagement.action')"
        align="center"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span class="link-primary action-link" @click="handleToggleEnv(scope.row)">{{ $t('page.deploy.deployClusters') }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"/>

    <el-dialog :visible.sync="dialogClusterVisible" :title="titleCluster">
      <el-table :data="viewCluster">
        <el-table-column :label="$t('page.deploy.clusterName')" property="date">
          <template slot-scope="scope">
            {{ scope.row.name }}
          </template>
        </el-table-column>
        <el-table-column :label="$t('page.deploy.clusterHosts')" property="name">
          <template slot-scope="scope">
            <div v-html="scope.row.hosts === '' ? '-': scope.row.hosts.split(',').join('<br>')"/>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog :title="$t('page.deploy.log')" :visible.sync="dialogJobLogVisible" width="80%" top="5vh">
      <div class="log" v-html="jobLog"/>
    </el-dialog>
  </div>
</template>

<script>
import apic from '@/api/const'
import apiDeployment from '@/api/deployment'
import apiJob from '@/api/job'
import apiJobStep from '@/api/jobStep'
import { Message } from 'element-ui'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ComplexTable',
  components: { Pagination },
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

      getRowKeys(row) {
        return row.deployment.id
      },
      expands: [],
      getJobRowKeys(row) {
        return row.job.id
      },
      jobExpands: [],
      titleCluster: '',
      dialogClusterVisible: false,
      viewCluster: [{
        name: '',
        hosts: ''
      }],

      // job detail steps
      jobDetailSteps: {
        // jobId => [{jobStep01}, {jobStep01}]
      },
      // job step's active index for element-ui step
      jobDetailStepsActive: {
        // jobId => 1
      },
      apic: apic,
      dialogJobLogVisible: false,
      jobLog: '',
      jobPrefix: 'job_prefix_'
    }
  },
  created() {
    this.getList()
  },
  mounted() {
  },
  methods: {
    getList() {
      const project = this.$store.state.project.curProject
      this.listLoading = true
      apiDeployment.listByProjectId(project.id).then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.total = response.data.total

        if (this.list.length > 0) {
          this.expands.push(this.list[0].deployment.id)
        }
      }).catch(error => {
        this.listLoading = false
        Message.error(error)
      })
    },

    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    handleStartDeploy(job) {
      this.$confirm(this.$t('page.deploy.tipStartDeploy'), this.$t('page.deploy.confirmDeploy'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiJob.start(job.id).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
        }).catch(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    },

    handleToggleEnv(row) {
      this.$refs.deployTable.toggleRowExpansion(row)
    },

    handleViewClusterHosts(row) {
      this.dialogClusterVisible = true
      this.titleCluster = row.cluster.name
      this.viewCluster = [{
        name: row.cluster.name,
        hosts: row.job.deploy_hosts
      }]
    },

    // handler toggle job steps
    handleToggleJobDetailSteps(row) {
      // tips: when row expansion changed, jobDetailStepsExpandChange(row) will be invoke
      this.$refs[this.jobPrefix + row.job.deployment_id].toggleRowExpansion(row)
    },
    // when job steps expands, this func will be invoke
    jobDetailStepsExpandChange(row) {
      this.getJobDetailSteps(row.job.id)
    },

    handleViewLog(job) {
      this.dialogJobLogVisible = true
      apiJobStep.getByJobId(job.id).then(response => {
        let log = ''
        response.data.list.forEach((item, k) => {
          log += '\n[' + (k + 1) + '. ' + this.$t('page.deploy.jobStepLabels.' + item.step) + '] \n' +
            item.cmd +
            '\n' +
            item.log
        })
        this.jobLog = log.split('\n').join('<br>')

        // check whether job ends
        if ([apic.JOB_STATUS_DEALING, apic.JOB_STATUS_READY].indexOf(job.status) !== -1) {
          this.jobLog += (this.jobLog !== '' ? '<br>' : '') + '<i class="el-icon-loading"></i>'
        }
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },

    getDeploymentStatusClass(status) {
      status = parseInt(status)
      switch (status) {
        case apic.DEPLOYMENT_STATUS_CREATED:
        case apic.DEPLOYMENT_STATUS_NO_AUDIT:
        case apic.DEPLOYMENT_STATUS_WAIT_AUDIT:
        case apic.DEPLOYMENT_STATUS_AUDIT_PASS:
        case apic.DEPLOYMENT_STATUS_DEALING:
          return 'status-process'
        case apic.DEPLOYMENT_STATUS_AUDIT_REJECT:
        case apic.DEPLOYMENT_STATUS_CANCELED:
          return 'status-cancel'
        case apic.DEPLOYMENT_STATUS_FAILED:
          return 'status-failed'
        case apic.DEPLOYMENT_STATUS_SUCC:
          return 'status-success'
        default:
          return ''
      }
    },
    getJobStatusClass(status) {
      status = parseInt(status)
      switch (status) {
        case apic.JOB_STATUS_CREATED:
        case apic.JOB_STATUS_READY:
        case apic.JOB_STATUS_DEALING:
          return 'status-process'
        case apic.JOB_STATUS_SUCC:
          return 'status-success'
        case apic.JOB_STATUS_FAILED:
          return 'status-failed'
        default:
          return ''
      }
    },
    formatJobShortSteps(status) {
      status = parseInt(status)
      const steps = [
        // ready
        {
          values: [apic.JOB_STATUS_CREATED, apic.JOB_STATUS_READY],
          success: apic.JOB_STATUS_READY,
          successLabel: '就绪',
          title: '',
          status: ''
        },
        // deploying
        {
          values: [apic.JOB_STATUS_DEALING],
          success: apic.JOB_STATUS_DEALING,
          successLabel: '部署命令',
          waitLabel: '部署',
          title: '',
          status: ''
        },
        // success or failed
        {
          values: [apic.JOB_STATUS_SUCC, apic.JOB_STATUS_FAILED],
          success: apic.JOB_STATUS_SUCC,
          successLabel: '成功',
          waitLabel: '完成',
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
            steps[i].title = this.$t('page.deploy.jobStatusLabels.' + status)
            if (apic.JOB_STATUS_FAILED === status) {
              steps[i].status = 'error'
            } else if (apic.JOB_STATUS_SUCC === status) {
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

      return steps
    },
    // request job step details
    getJobDetailSteps(jobId) {
      // init default job step details
      if (!this.jobDetailSteps[jobId]) {
        this.$set(this.jobDetailSteps, jobId, this.formatJobDetailStepsByList([]))
      }

      // request job steps
      apiJobStep.getByJobId(jobId).then(response => {
        // this.jobDetailSteps[jobId] = this.formatJobDetailStepsByList(response.data.list)
        this.$set(this.jobDetailSteps, jobId, this.formatJobDetailStepsByList(response.data.list))
        this.$set(this.jobDetailStepsActive, jobId, response.data.list.length)
      }).catch(error => {
        console.log(error)
        // Message.error(error)
      })
    },

    // format job steps for render by job steps
    formatJobDetailStepsByList(jobSteps) {
      const list = []
      let curJobStep = 0
      if (jobSteps.length > 0) {
        curJobStep = jobSteps[jobSteps.length - 1].step
      }

      const steps = [
        apic.JOB_STEP_INIT_WORKSPACE,
        apic.JOB_STEP_SYNC_VERSION_PACKAGE,
        apic.JOB_STEP_UNPACK_VERSION_PACKAGE,
        apic.JOB_STEP_CMD_BEFORE_DEPLOY,
        apic.JOB_STEP_DO_DEPLOY,
        apic.JOB_STEP_CMD_AFTER_DEPLOY,
        apic.JOB_STEP_CMD_HEALTH_CHECK,
        apic.JOB_STEP_CMD_ONLINE,
        apic.JOB_STEP_END_CLEAN
      ]
      steps.forEach((v, k) => {
        const tmp = {
          step: v,
          label: this.$t('page.deploy.jobStepLabels.' + v),
          status: 'wait'
        }
        if (curJobStep < v) {
          tmp.status = 'wait'
        } else if (curJobStep === v) {
          if (jobSteps[k].status === apic.JOB_STEP_STATUS_FAILED) { // job step failed
            tmp.status = 'fail'
          } else if (jobSteps[k].status === apic.JOB_STEP_STATUS_SUCC) { // job step success
            tmp.status = 'success'
          } else { // job step processing
            tmp.status = 'process'
          }
        } else {
          tmp.status = 'success'
        }

        list.push(tmp)
      })

      return list
    },

    getJobActive(jobStatus) {
      switch (jobStatus) {
        case apic.JOB_STATUS_CREATED:
          return 0
        case apic.JOB_STATUS_READY:
          return 1
        case apic.JOB_STATUS_DEALING:
          return 2
        case apic.JOB_STATUS_FAILED:
        case apic.JOB_STATUS_SUCC:
          return 3
        default:
          return 0
      }
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
  .status-cancel{
    color: #999;
  }
  .log {
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
