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
                
                <li v-show="role == 'admin'">
                    <router-link
                        :to="{ name: 'adminGetCategories'}">
                            <a href="#"><i class="fa fa-fw fa-group"></i> Категорії</a>
                    </router-link> 
                </li>

                <li v-show="role == 'admin'">
                    <router-link
                        :to="{ name: 'adminGetRegions'}">
                            <a href="#"><i class="fa fa-fw fa-globe"></i> Регіони</a>
                    </router-link> 
                </li>

                <li v-show="role == 'admin' || role == 'owner'">
                    <!-- <a href="/admin/tours"><i class="fa fa-desktop"></i> Tours</a> -->
                    <router-link
                        :to="{ name: 'adminGetTours'}">
                            <a href="#"><i class="fa fa-fw fa-plane"></i> Тури</a>
                    </router-link> 
                </li>

                <li v-show="role == 'admin'">
                    <router-link
                        :to="{ name: 'adminGetEvents'}">
                            <a href="#"><i class="fa fa-compass"></i> Події</a>
                    </router-link>
                </li>
                <li v-show="role == 'admin' || role == 'owner'">
                    <!-- <a href="/admin/tours"><i class="fa fa-desktop"></i> Tours</a> -->
                    <router-link
                        :to="{ name: 'adminGetTrips'}">
                            <a href="#"><i class="fa fa-car"></i> Поїздки</a>
                    </router-link> 
                </li>
                <li>
                    <router-link
                        :to="{ name: 'adminGetOrders'}">
                            <a href="#"><i class="fa fa-file-text-o"></i> Замовлення</a>
                    </router-link> 
                </li>
                <li>
                    <router-link
                        :to="{ name: 'adminGetComments'}">
                            <a href="#"><i class="fa fa-fw fa-comment"></i> Коментарі</a>
                    </router-link> 
                </li>
                <li v-show="role == 'client' || role == 'owner'">
                    <router-link
                        :to="{ name: 'adminEditProfile'}">
                            <a href="#"><i class="fa fa-edit"></i> Редагувати профіль</a>
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