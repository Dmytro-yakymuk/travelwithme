<template>
    <nav class="navbar-default navbar-side" role="navigation">
        <div class="sidebar-collapse">
            <ul class="nav" id="main-menu">

                <li>
                    <!-- <a class="active-menu" href="/admin"><i class="fa fa-dashboard"></i> Dashboard</a> -->
                    <router-link
                        :to="{ name: 'admin'}">
                            <a href="#"><i class="fa fa-dashboard"></i> Головна</a>
                    </router-link> 
                </li>
                
                <li v-show="role == 'admin' || role == 'owner'">
                    <!-- <a href="/admin/tours"><i class="fa fa-desktop"></i> Tours</a> -->
                    <router-link
                        :to="{ name: 'adminGetTours'}">
                            <a href="#"><i class="fa fa-desktop"></i> Тури</a>
                    </router-link> 
                </li>

            </ul>

        </div>

    </nav>
</template>

<script>
export default {
    name: "AdminSidebar",
    data: function () {
        return {
            role: localStorage.getItem('role'),
        }
    },
}
</script>

<style scoped>

</style>