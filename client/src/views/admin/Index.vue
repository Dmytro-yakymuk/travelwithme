<template>
    <div id="page-wrapper">

        <h1><small>Вітаю {{ role == 'admin' ? 'Адмін' : oneUser.name }}!</small></h1>
        <br>

        <div v-if="role == 'admin'" class="col-md-4 col-sm-12 col-xs-12">
            <h3 class="text_color">Загальні дані по системі:</h3>

            <a href="#" class="list-group-item">
                <span class="badge">23</span>
                <i class="fa fa-fw fa-group"></i> Категорій
            </a>

            <a href="#" class="list-group-item">
                <span class="badge">3</span>
                <i class="fa fa-fw fa-globe"></i>Регіонів
            </a>

            <a href="#" class="list-group-item">
                <span class="badge">36</span>
                <i class="fa fa-fw fa-plane"></i> Турів
            </a>
            <a href="#" class="list-group-item">
                <span class="badge">7</span>
                <i class="fa fa-fw fa-comment"></i> Коментарів
            </a>
            <a href="#" class="list-group-item">
                <span class="badge">36</span>
                <i class="fa fa-file-text-o"></i> Замовлень
            </a>
            <a href="#" class="list-group-item">
                <span class="badge">111</span>
                <i class="fa fa-car"></i> Поїздок
            </a>
        </div>

        <div v-if="role == 'client' && oneUser != null || role == 'owner' && oneUser != null" class="col-md-4 col-sm-12 col-xs-12">
            <h3 class="text_color">Мої дані:</h3>

            <a href="#" class="list-group-item">
                <span class="badge">{{ oneUser.email }}</span>
                <i class="fa fa-fw fa-envelope"></i> Email
            </a>
            <a href="#" class="list-group-item">
                <span class="badge">{{ oneUser.phone }}</span>
                <i class="fa fa-fw fa-phone"></i> Номер телефону
            </a>
        </div>

    </div>
</template>

<script>
import {mapGetters, mapActions} from "vuex"
export default {
    name: "AdminIndex",
    data: function () {
        return {
            role: localStorage.getItem('role'),
        }
    },
    computed: mapGetters(['oneUser']),
    methods: {
        ...mapActions(['getOneUser']),
    },
    async mounted() {
        this.getOneUser();
    }
}
</script>

<style scoped>
    .text_color {
        color: #777
    }
    .mar-tp {
        margin-top: 76px;
    }
</style>
