<template>
  <el-scrollbar wrap-class="scrollbar-wrapper">
    <el-menu
      :show-timeout="200"
      :default-active="$route.path"
      :collapse="isCollapse"
      mode="vertical"
      background-color="#304156"
      text-color="#fff"
      active-text-color="#409EFF"
    >
      <div class="menu-wrapper" style="border-bottom: 1px dashed rgb(42,57,75); padding-bottom: 5px; margin-bottom: 5px;margin-top: 5px">
        <!--<a href="#/deploy/index" class="">-->
        <li
          role="menuitem"
          tabindex="-1"
          class="el-menu-item submenu-title-noDropdown"
          style="padding-left: 12px; color: white; background-color: transparent;">
          <i class="el-icon-menu"/>
          <span>{{ project.name }}</span>
        </li>
        <!--</a>-->
      </div>

      <sidebar-item v-for="route in permission_routers" :key="route.path" :item="route" :base-path="route.path"/>
    </el-menu>
  </el-scrollbar>
</template>

<script>
import { mapGetters } from 'vuex'
import SidebarItem from './SidebarItem'

export default {
  components: { SidebarItem },
  computed: {
    ...mapGetters([
      'permission_routers',
      'sidebar'
    ]),
    isCollapse() {
      return !this.sidebar.opened
    },
    project() {
      return this.$store.state.project.curProject
    }
  }
}
</script>
