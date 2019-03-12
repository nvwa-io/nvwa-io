<template>
  <div class="app-container">
    <div class="table-out-title"> {{ $t('page.admin.deploy.manageUser') }}<span class="tail"/></div>

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
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.admin.deploy.displayName')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.display_name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.admin.deploy.account')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.admin.deploy.email')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.email }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.admin.deploy.systemRole')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.role }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('page.deploy.utime')" min-width="90" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.utime }}</span>
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
  </div>

</template>

<script>
import apiUser from '@/api/admin/user'
import { Message } from 'element-ui'
// import waves from '@/directive/waves' // Waves directive
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ComplexTable',
  components: { Pagination },
  // directives: { waves },
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
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      apiUser.list().then(response => {
        this.listLoading = false
        this.list = response.data.list
        this.total = response.data.total
      }).catch(error => {
        this.listLoading = false
        Message.error(error)
      })
    },

    handleUpdate(row) {
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .CodeMirror-gutter-wrapper {
    left: 35px!important;
  }
</style>
