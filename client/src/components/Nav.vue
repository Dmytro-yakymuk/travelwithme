<template>
    <!-- main header -->
    <header class="white-header">
        <div class="container-fluid">
            <!-- logo -->
            <div class="logo">
                <a href="/">
                   <h1>TravelWithMe</h1>
                </a>
            </div>
            <!-- main navigation -->
            <nav class="navbar navbar-default">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle nav-opener" data-toggle="collapse" data-target="#nav">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                </div>
                <!-- main menu items and drop for mobile -->
                <div class="collapse navbar-collapse" id="nav">
                    <!-- main navbar -->
                    <ul class="nav navbar-nav">
                        <li class="dropdown">
                            <a href="/">Головна <b class="icon-angle-down"></b></a>
                        </li>
                        <li class="dropdown">
                            <a href="/tours">Каталог <b class="icon-angle-down"></b></a>
                        </li>
                        
                        
                        
                        <li v-if="!loggedIn" class="visible-xs visible-sm">
                            <a href="/login">
                                <span class="icon icon-user"></span>
                                <span class="text">Вхід</span>
                            </a>
                        </li>

                        <li v-if="!loggedIn" class="hidden-xs hidden-sm v-divider">
                            <a href="/login">
                                <span class="icon icon-user"></span>
                            </a>
                        </li>

                        <li v-if="loggedIn" class="hidden-xs hidden-sm v-divider">
                            <a href="/admin">
                                <span class="icon icon-user"></span>
                            </a>
                        </li>

                        <li v-if="loggedIn" class="hidden-xs hidden-sm v-divider">
                            <a href="/logout">
                                <span class="fa fa-sign-out fa-fw"></span>
                            </a>
                        </li>   

                        <li class="visible-xs visible-sm nav-visible dropdown last-dropdown v-divider">
                            <a href="/cart" data-toggle="dropdown">
                                <span class="icon icon-cart"></span>
                                <span class="text hidden-md hidden-lg">Cart</span>
                            </a>
                            <div v-if="allTripsWhichAttach != null && allTripsWhichAttach.length != 0" class="dropdown-menu dropdown-md">
                                <div class="drop-wrap cart-wrap">
                                    <strong class="title">Кошик для покупок</strong>
                                    <ul class="cart-list">
                                        <li v-for="trip in allTripsWhichAttach" :key="trip.id">
                                            <div class="img">
                                                <a :href="'/tour/'+trip.slug">
                                                    <img
                                                        :src="imageURL+trip.image" 
                                                        height="165" width="170" 
                                                        alt="image description">
                                                </a>
                                            </div>
                                            <div class="text-holder">
                                                <span class="amount">x {{ trip.count }}</span>
                                                <div class="text-wrap">
                                                    <strong class="name"><a :href="'/tour/'+trip.slug">{{ trip.title }}</a></strong>
                                                    <span class="price">{{ trip.price * trip.count }} грн</span>
                                                </div>
                                            </div>
                                        </li>
                                        
                                    </ul>
                                    <div class="footer">                        
                                        <a href="/cart" class="btn btn-primary">Детальніше</a>
                                        <span class="total">{{ getTotalPrice }} грн</span>
                                    </div>
                                </div>
                            </div>
                        </li>
                        
                    </ul>
                </div>
            </nav>
        </div>
    </header>
    <!-- main banner -->
</template>

<script>

    import {mapGetters, mapActions} from "vuex"
    import {ShowImageURL} from '../constants/index'

    export default {
        name:"Nav",
        data: function () {
			return {
				imageURL: ShowImageURL,
			}
		},
        computed: mapGetters(['allTripsWhichAttach', 'getTotalPrice', 'loggedIn']),
        methods: {
            ...mapActions(['getAllTripsWhichAttach']),
        },
        async mounted() {
            this.getAllTripsWhichAttach();
        }
    }
</script>

<style scoped>
    .white-header{
        background: #252525!important;
        border-bottom: 1px solid #181818;
    }

    .logo {
        border-right: 0px;
    }

</style>