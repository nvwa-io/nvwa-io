<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" size="mini" icon="el-icon-plus" type="primary" @click="handleCreate">
        {{ $t('page.admin.deploy.addProjectRole') }}
      </el-button>
    </div>
    <div class="table-out-title"> {{ $t('page.admin.deploy.manageProjectRole') }}<span class="tail"/></div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="projectRoleList"
      border
      fit
      stripe
      style="width: 100%;"
      class="table-primary">
      <el-table-column :label="$t('page.deploy.id')" prop="id" align="center" width="45">
        <template slot-scope="scope">
          <span>{{ scope.row.project_role.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.admin.deploy.projectRoleName')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.project_role.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.utime')" min-width="90" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.project_role.utime }}</span>
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

    <el-dialog :title="textRoleMap[dialogRoleStatus]" :visible.sync="dialogRoleVisible" width="500px" top="8vh">
      <el-form
        ref="dataForm"
        :model="formRoleData"
        label-position="left"
        label-width="70px"
        style="margin: auto 20px">

        <el-row :gutter="30">
          <div style="margin-bottom: 15px">
            <el-input v-model="formRoleData.projectRoleName" :placeholder="$t('page.admin.deploy.projectRoleName')" :value="formRoleData.projectRoleName"/>
          </div>

          <el-table
            :data="permScopeList"
            style="width: 100%"
            class="table-primary">
            <el-table-column :label="$t('page.deploy.permScope')" align="left" width="120px">
              <template slot-scope="scope">
                {{ $t('page.deploy.' + scope.row.key) }}
              </template>
            </el-table-column>
            <el-table-column :label="$t('page.deploy.permItem')">
              <template slot-scope="scope">
                <el-checkbox-group
                  v-model="formRoleData.checkedPerms"
                  :min="1"
                >
                  <el-checkbox v-for="item in scope.row.list" :label="item" :key="item" >{{ $t('page.deploy.' + item.split('.').join('___')) }}</el-checkbox>
                </el-checkbox-group>
              </template>
            </el-table-column>
          </el-table>

        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogRoleVisible= false">{{ $t('table.cancel') }}</el-button>
        <el-button :loading="loadingRoleForm" type="primary" @click="dialogRoleStatus=== 'create'? handleAddRoleSubmit(): handleEditRoleSubmit()">{{ $t('table.confirm') }}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { Message } from 'element-ui'
import apiProjectRole from '@/api/projectRole'
import apiAdminProjectRolePerm from '@/api/admin/projectRolePerm'

export default {
  name: 'AdminProjectRole',
  data() {
    return {
      tableKey: 0,
      projectRoleList: null,
      listLoading: true,

      textRoleMap: {
        create: this.$t('page.admin.deploy.addProjectRole'),
        update: this.$t('page.admin.deploy.editProjectRole')
      },
      dialogRoleStatus: 'create',
      dialogRoleVisible: false,

      loadingRoleForm: false,

      rolePerms: [
        // Permission01
      ],
      projectRoleIdMapPerm: {
        // project_role_id => [{Permission}]
      },

      // role from
      formRoleData: {
        id: undefined,
        projectRoleName: '',
        checkedPerms: []
      },

      // all project role permissions
      allRolePerms: [
        // {key:'xxx',label:'xxx'}
      ],

      // gather scope permissions
      // e.g {'app' => ["app.create","app.delete","app.update"]}
      permScopeMap: {},
      // trans this.permScopeMap to permScopeList
      // [{key: 'app',permItem: ["app.create", "app.delete"]}]
      permScopeList: []
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true

      apiProjectRole.list().then(response => {
        this.listLoading = false
        this.projectRoleList = response.data.list
        this.projectRoleList.forEach((item, k) => {
          this.projectRoleIdMapPerm[item.project_role.id] = item.project_role_perm
          // if (callback.length > 0) {
          //   callback[0]()
          // }
        })
        console.log(this.projectRoleList)
      }).catch(error => {
        this.listLoading = false
        console.log(error)
        Message.error(error)
      })
    },

    // gather scope permissions
    getAllScopePerm() {
      this.permScopeMap = {}
      this.permScopeList = []

      return apiAdminProjectRolePerm.all().then(response => {
        // gather scope permissions
        // e.g [{'app' => ["app.create","app.delete","app.update"]}]
        const tmpList = response.data.list
        const perms = Object.getOwnPropertyNames(tmpList)
        // gather same scope, e.g: app, project's permission
        // according to {xxx}.xxxx
        perms.forEach((perm, k) => {
          const arr = perm.split('.')
          if (arr.length < 2) {
            console.log('invalid permission: ' + perm)
            return
          }

          const key = arr[0]
          if (!this.permScopeMap[key]) {
            this.permScopeMap[key] = []
          }
          this.permScopeMap[key].push(perm)
        })

        Object.getOwnPropertyNames(this.permScopeMap).forEach((key, i) => {
          if (key === '__ob__') {
            return
          }
          this.permScopeList.push({
            key: key,
            list: this.permScopeMap[key]
          })
        })
      }).catch(error => {
        console.log(error)
      })
    },

    handleCreate() {
      this.dialogRoleVisible = true
      this.dialogRoleStatus = 'create'

      this.formRoleData.id = undefined
      this.formRoleData.projectRoleName = ''
      this.formRoleData.checkedPerms = []

      this.getAllScopePerm().then(() => {
        this.permScopeList.forEach((item, v) => {
          item.list.forEach((perm, i) => {
            this.formRoleData.checkedPerms.push(perm)
          })
        })
      })
    },

    handleUpdate(row) {
      this.dialogRoleStatus = 'update'
      this.dialogRoleVisible = true

      // set project_role.id
      this.formRoleData.id = row.project_role.id
      this.formRoleData.projectRoleName = row.project_role.name
      this.getAllScopePerm().then(() => {
        this.formRoleData.checkedPerms = []
        row.project_role_perm.forEach((item, i) => {
          this.formRoleData.checkedPerms.push(item.perm)
        })
      })
    },

    // add project role submitted
    handleAddRoleSubmit() {
      this.loadingRoleForm = true
      apiAdminProjectRolePerm.batchCreate(this.formRoleData.projectRoleName, this.formRoleData.checkedPerms)
        .then(response => {
          this.loadingRoleForm = false
          this.dialogRoleVisible = false
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
        }).catch(error => {
          this.loadingRoleForm = false
          console.log(error)
        })
    },

    // edit project role submitted
    handleEditRoleSubmit() {
      this.loadingRoleForm = true
      apiAdminProjectRolePerm.batchUpdate(this.formRoleData.id, this.formRoleData.projectRoleName, this.formRoleData.checkedPerms)
        .then(response => {
          this.loadingRoleForm = false
          this.dialogRoleVisible = false
          Message.success(this.$t('page.deploy.operationSuccess'))
          this.getList()
        }).catch(error => {
          this.loadingRoleForm = false
          console.log(error)
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
