<template>
    <nav class="navbar navbar-default top-navbar" role="navigation">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".sidebar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <router-link :to="{ name: 'admin'}"><a class="navbar-brand" href="#">Адмінка:)</a></router-link>   
        </div>

        <ul class="nav navbar-top-links navbar-right">
            <li><a href="/">Тур агенство</a></li>

            <li class="divider"></li>
            <li>
                <router-link :to="{ name: 'logout'}">
                    <a href="#"><i class="fa fa-sign-out fa-fw"></i> Вихід</a>
                </router-link> 
            </li>
           
        </ul>
    </nav>
</template>

<script>
  export default {
    name:"AdminNav"
  }
</script>

<style scoped>

</style>