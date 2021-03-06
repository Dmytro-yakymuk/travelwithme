<template>
  <div id="app">
    <!-- main wrapper -->
    <div id="wrapper">
      <div class="page-wrapper">
        <Nav />
        <Banner v-show="$route.name=='home'" />
        <router-view/>
      </div>
      <Footer />
    </div>

    <!-- scroll to top -->
    <div class="scroll-holder text-center">
      <a href="javascript:" id="scroll-to-top"><i class="icon-arrow-down"></i></a>
    </div>

  </div>
</template>

<style scoped>
	@import "../../assets/vendors/font-awesome/css/font-awesome.css";
	@import "../../assets/vendors/material-design-icons/material-icons.css";
	@import "../../assets/css/fonts/icomoon/icomoon.css"; 
	@import "../../assets/css/bootstrap.css";
	@import "../../assets/css/main.css";
</style>

<script>
  import Nav from "../../components/Nav"
  import Footer from "../../components/Footer"
  import Banner from "../../components/Banner"
  export default {
    name: 'DefaultLayout',
    components: {Nav, Footer, Banner}
  }
</script>
