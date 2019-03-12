<template>
  <div>
    <home-header/>
    <router-view/>
    <home-footer/>
  </div>
</template>

<script>
import HomeHeader from '@/views/home/components/Header'
import HomeFooter from '@/views/home/components/Footer'

export default {
  name: 'HomeLayout',
  components: {
    HomeHeader,
    HomeFooter
  },
  data() {
    return {

    }
  },
  methods: {

  }
}
</script>

