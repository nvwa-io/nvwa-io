<template>
  <div>
    <div class="dashboard-editor-container">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/home' }">{{ $t('page.deploy.home') }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ $t('page.deploy.projectManagement') }}</el-breadcrumb-item>
      </el-breadcrumb>
      <br>
      <br>

      <div>
        <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleCreateProject">
          {{ $t('page.deploy.addProject') }}
        </el-button>
      </div>
      <br>
      <!--<div class="table-out-title blk"> {{ $t('page.deploy.myProject') }} <span class="tail"/></div>-->
      <!--
        :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
      -->
      <el-table
        v-loading="loadingProjects"
        :data="projects"
        border
        fit
        stripe
        style="width: 100%"
        class="table-primary">
        <el-table-column
          :label="$t('page.deploy.projectName')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.project.name }}
          </template>

        </el-table-column>

        <el-table-column
          :label="$t('page.deploy.description')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.project.description }}
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('page.deploy.creator')"
          align="center">
          <template slot-scope="scope">
            {{ scope.row.user.username }}
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
          :label="$t('page.deploy.memberManagement')"
          align="center">
          <template slot-scope="scope">
            <span class="link-primary action-link" @click="handleMemberManagement(scope.row.project.id)">{{ $t('page.deploy.member') }}</span>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('page.action')"
          align="center">
          <template slot-scope="scope">
            <span class="link-primary action-link" @click="handleEditProject(scope.row.project)">{{ $t('page.opTable.edit') }}</span>
            <span class="link-danger action-link" @click="handleDeleteProject(scope.row.project)">{{ $t('page.opTable.delete') }}</span>
          </template>
        </el-table-column>
      </el-table>

      <div style="text-align: center;color: #7d8899">
        <br>
        <br>
        <br>
        <p> {{ $t('page.deploy.nvwaArch') }}</p>
        <br>
        <div>
          <img src="@/assets/images/nvwa-arch.png" alt="nvwa-arch" style="width:100%;">
        </div>
      </div>

      <el-dialog :visible.sync="dialogMemberVisible" :title="$t('page.deploy.memberManagement')" width="700px">
        <div style="padding-bottom:20px">
          <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleCreateMember">
            {{ $t('page.deploy.addMember') }}
          </el-button>
        </div>

        <el-table
          :data="memberList"
          border
          fit
          stripe
          style="width: 100%"
          class="table-primary">
          <el-table-column :label="$t('page.deploy.username')" align="center">
            <template slot-scope="scope">
              {{ scope.row.user.username }}
            </template>
          </el-table-column>
          <el-table-column :label="$t('page.deploy.role')" align="center">
            <template slot-scope="scope">
              {{ scope.row.project_role.name }}
            </template>
          </el-table-column>
          <el-table-column :label="$t('page.deploy.ctime')" align="center">
            <template slot-scope="scope">
              {{ scope.row.project_role.ctime }}
            </template>
          </el-table-column>
          <el-table-column
            :label="$t('page.action')"
            align="center">
            <template slot-scope="scope">
              <span class="link-primary action-link" @click="handleEditRole(scope.row)">{{ $t('page.deploy.editRole') }}</span>
              <span class="link-danger action-link" @click="handleRemoveMember(scope.row.member)">{{ $t('page.deploy.remove') }}</span>
            </template>
          </el-table-column>
        </el-table>
        <div style="height: 35px;"/>
      </el-dialog>
    </div>

    <el-dialog :title="textMemberMap[dialogMemberStatus]" :visible.sync="dialogFormMemberVisible" width="700px" top="8vh">
      <el-form
        ref="dataForm"
        :model="formMemberData"
        label-position="left"
        label-width="70px"
        style="margin: auto 20px">

        <el-row :gutter="30">
          <el-col :xs="10" :sm="10" :lg="10">
            <el-form-item v-if="dialogMemberStatus === 'update'" :label="$t('page.deploy.username')" prop="username">
              <el-input v-model="formMemberData.username" :value="formMemberData.username" disabled/>
            </el-form-item>

            <el-form-item v-if="dialogMemberStatus === 'create'" :label="$t('page.deploy.username')">
              <el-select
                v-model="formMemberData.uid"
                :loading="loadingAllUsers"
                :loading-text="$t('page.deploy.loadingUsers')"
                :placeholder="$t('page.deploy.selectUser')"
                filterable
                class="filter-item"
                style="width: 100%;">
                <el-option v-for="item in allUsers" :key="item.id" :label="item.display_name" :value="item.id"/>
              </el-select>
            </el-form-item>

            <el-form-item :label="$t('page.deploy.role')">
              <el-select
                v-model="formMemberData.project_role_id"
                :loading="loadingRole"
                :loading-text="$t('page.deploy.loadingRole')"
                :placeholder="$t('page.deploy.selectRole')"
                filterable
                class="filter-item"
                style="width: 100%;"
                @change="changeSelectRole">
                <el-option v-for="item in projectRoleList" :key="item.project_role.id" :label="item.project_role.name" :value="item.project_role.id"/>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="14" :sm="14" :lg="14">
            <el-table
              :data="gatherRolePermission(rolePerms)"
              style="width: 100%"
              class="table-primary">
              <el-table-column :label="$t('page.deploy.permScope')" align="left" width="120px">
                <template slot-scope="scope">
                  {{ $t('page.deploy.' + scope.row.key) }}
                </template>
              </el-table-column>
              <el-table-column :label="$t('page.deploy.permItem')">
                <template slot-scope="scope">
                  <el-tag v-for="item in scope.row.list" :key="item.id" class="perm-item" size="mini">{{ $t('page.deploy.' + item.perm.split('.').join('___')) }}</el-tag>
                </template>
              </el-table-column>
            </el-table>
          </el-col>

        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormMemberVisible= false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingMemberForm" type="primary" @click="dialogMemberStatus === 'create'? handleAddMemberSubmit(): handleEditRoleSubmit()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>

    <!--project form-->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="550px">
      <el-form
        ref="dataForm"
        :rules="projectRules"
        :model="formProjectData"
        label-position="left"
        label-width="90px"
        style="padding: 0 20px">
        <el-form-item :label="$t('page.deploy.projectName')" prop="name">
          <el-input v-model="formProjectData.name" placeholder="e.g: 女娲项目" />
        </el-form-item>
        <el-form-item :label="$t('page.deploy.description')" prop="description">
          <el-input
            :autosize="{ minRows: 2, maxRows: 4}"
            v-model="formProjectData.description"
            type="textarea"
            placeholder="e.g: 开源 DevOps 部署系统"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingProjectForm" type="primary" @click="dialogStatus==='create'?createProject():updateProject()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>

import apiUser from '@/api/user'
import apiProject from '@/api/project'
import apiMember from '@/api/member'
import apiProjectRole from '@/api/projectRole'
import { Message } from 'element-ui'

export default {
  name: 'ProjectManagement',
  components: {},
  data() {
    return {
      projects: [],
      loadingProjects: false,
      search: '',
      dialogMemberVisible: false,
      memberList: [],

      // project form
      formProjectData: {
        id: undefined,
        name: '',
        description: ''
      },
      projectRules: {
        name: [{ required: true, message: '项目名称不能为空', trigger: 'blur' }]
      },
      loadingProjectForm: false,
      dialogStatus: 'create',
      textMap: {
        create: this.$t('page.deploy.addProject'),
        update: this.$t('page.deploy.editProject')
      },
      dialogFormVisible: false,

      // member from
      formMemberData: {
        id: undefined,
        username: '',
        uid: undefined,
        project_id: 0,
        project_role_id: undefined
      },
      projectRoleList: [],
      loadingAllUsers: false,
      loadingRole: false,
      rolePerms: [
        // Permission01
      ],
      projectRoleIdMapPerm: {
        // project_role_id => [{Permission}]
      },
      loadingMemberForm: false,
      dialogMemberStatus: 'create',
      textMemberMap: {
        create: this.$t('page.deploy.addMember'),
        update: this.$t('page.deploy.editRole')
      },
      dialogFormMemberVisible: false,
      allUsers: []

    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loadingProjects = true
      apiProject.list().then(response => {
        this.loadingProjects = false
        this.projects = response.data.list
      }).catch(error => {
        this.loadingProjects = false
        console.log(error)
        Message.error(error)
      })
    },

    handleCreateProject() {
      this.resetFormProjectData()
      this.dialogFormVisible = true
    },
    // create project submit
    createProject() {
      this.loadingProjectForm = true
      apiProject.create(this.formProjectData.name, this.formProjectData.description).then(response => {
        this.loadingProjectForm = false
        this.dialogFormVisible = false
        Message.success(this.$t('page.deploy.operationSuccess'))
        this.getList()
      }).catch(error => {
        this.loadingProjectForm = false
        console.log(error)
        Message.error(error)
      })
    },

    resetFormProjectData() {
      this.formProjectData.id = undefined
      this.formProjectData.name = ''
      this.formProjectData.description = ''
    },

    handleEditProject(project) {
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.resetFormProjectData()

      this.formProjectData.id = project.id
      this.formProjectData.name = project.name
      this.formProjectData.description = project.description
    },
    // update project submit
    updateProject() {
      this.loadingProjectForm = true
      apiProject.update(this.formProjectData.id, this.formProjectData.name, this.formProjectData.description).then(response => {
        this.loadingProjectForm = false
        this.dialogFormVisible = false
        Message.success(this.$t('page.deploy.operationSuccess'))
        this.getList()
      }).catch(error => {
        this.loadingProjectForm = false
        console.log(error)
        Message.error(error)
      })
    },

    // delete project
    handleDeleteProject(project) {
      this.$confirm(this.$t('page.deploy.tipDeleteProject'), this.$t('page.deploy.deleteConfirm'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiProject.deleteById(project.id).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
        }).then(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    },

    // reset form member
    resetFormMemberData() {
      this.formMemberData.id = undefined
      this.formMemberData.username = ''
      this.formMemberData.uid = undefined
      this.formMemberData.project_role_id = undefined
      // don't reset project id
    },

    // show create member form dialog
    handleCreateMember() {
      this.dialogMemberStatus = 'create'
      this.resetFormMemberData()
      this.dialogFormMemberVisible = true

      // request all users
      this.getAllUsers()

      // request all project roles
      this.getProjectRoles()
    },

    getAllUsers() {
      this.loadingAllUsers = true
      apiUser.all().then(response => {
        this.loadingAllUsers = false
        this.allUsers = response.data.list
      }).catch(error => {
        this.loadingAllUsers = false
        console.log(error)
        Message.error(error)
      })
    },

    handleMemberManagement(projectId) {
      this.dialogMemberVisible = true
      this.formMemberData.project_id = projectId
      this.getProjectMembers(projectId)
    },
    getProjectMembers(projectId) {
      apiMember.getListByProjectId(projectId).then(response => {
        this.memberList = response.data.list
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },
    handleEditRole(row) {
      this.dialogFormMemberVisible = true
      this.dialogMemberStatus = 'update'
      this.formMemberData.username = row.user.username
      this.formMemberData.id = row.member.id
      this.formMemberData.project_id = row.member.project_id

      // get project roles
      this.getProjectRoles(() => {
        // init current project_role_id
        this.formMemberData.project_role_id = row.project_role.id
        this.changeSelectRole(row.project_role.id)
      })
    },

    // get project roles and int role permission items
    getProjectRoles(...callback) {
      apiProjectRole.list().then(response => {
        this.projectRoleList = response.data.list
        this.projectRoleList.forEach((item, k) => {
          this.projectRoleIdMapPerm[item.project_role.id] = item.project_role_perm
          if (callback.length > 0) {
            callback[0]()
          }
        })
      }).catch(error => {
        console.log(error)
        Message.error(error)
      })
    },

    // invoke this func when member's project role id select (or set value)
    changeSelectRole(projectRoleId) {
      this.rolePerms = this.projectRoleIdMapPerm[projectRoleId]
    },

    // gather role permission
    gatherRolePermission(rolePerms) {
      if (!rolePerms) {
        return
      }

      const map = {}
      const keys = []
      const list = [] // {key:[{Permission}]}
      rolePerms.forEach((item, k) => {
        const arr = item.perm.split('.')
        if (arr.length < 2) {
          console.log('invalid permission: ' + item.perm)
          return
        }

        const key = arr[0]
        if (!map[key]) {
          map[key] = []
        }
        keys.push(key)
        map[key].push(item)
      })

      Object.getOwnPropertyNames(map).forEach((key, i) => {
        list.push({
          key: key,
          list: map[key]
        })
      })

      return list
    },

    // submit create member
    handleAddMemberSubmit() {
      this.loadingMemberForm = true
      console.log(this.formMemberData)
      apiMember.create({
        project_id: this.formMemberData.project_id,
        uid: this.formMemberData.uid,
        project_role_id: this.formMemberData.project_role_id
      }).then(response => {
        this.loadingMemberForm = false
        this.dialogFormMemberVisible = false
        Message.success(this.$t('page.deploy.operationSuccess'))
        this.getProjectMembers(this.formMemberData.project_id)
      }).catch(error => {
        this.loadingMemberForm = false
        console.log(error)
        Message.error(error)
      })
    },

    // submit edit role
    handleEditRoleSubmit() {
      this.loadingMemberForm = true
      apiMember.editRole(this.formMemberData.id, this.formMemberData.project_role_id).then(response => {
        this.loadingMemberForm = false
        this.dialogFormMemberVisible = false
        Message.success(this.$t('page.deploy.operationSuccess'))
        this.getProjectMembers(this.formMemberData.project_id)
      }).catch(error => {
        this.loadingMemberForm = false
        console.log(error)
        Message.error(error)
      })
    },
    handleRemoveMember(member) {
      this.$confirm(this.$t('page.deploy.tipRemoveMember'), this.$t('page.deploy.titleRemoveMember'), {
        confirmButtonText: this.$t('page.deploy.confirm'),
        cancelButtonText: this.$t('page.deploy.cancel')
      }).then(() => {
        apiMember.removeMember(member.id).then(response => {
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getProjectMembers(member.project_id)
        }).catch(error => {
          console.log(error)
          Message.error(error)
        })
      }).catch(() => {})
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  margin: 0 auto;
  width: 80%;
  min-width: 800px;
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
</style>
