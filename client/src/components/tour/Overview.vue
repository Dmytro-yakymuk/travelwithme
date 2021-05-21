<template>
    <div class="tab-pane active" id="tab1">
        <div class="row">
            <div class="col-md-6">
                <strong class="header-box">Детальніше про тур</strong>
                <div v-if="oneTour.text != ''" class="detail">
                    {{ oneTour.text }}
                </div>
            </div>
            <div class="col-md-6">
                <strong class="header-box">Огляд включень в туристичний пакет</strong>
                <div class="text-box" v-if="oneTour.events && oneTour.events.length > 0">
                    <div class="holder">
                        <strong class="title">Що входить у цей тур</strong>
                        <ul class="content-list tick-list">
                            <li v-for="event in oneTour.events" :key="event.id" >
                                <p>{{ event.text }}</p>
                            </li>
                        </ul>
                    </div>
                </div>
                <div v-else class="text-box not-included">
                    <div class="holder">
                        <strong class="title">Нічого не входить у цей тур</strong>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>   
    import {mapGetters} from "vuex"

    export default {
        name: 'overview',
        computed: mapGetters(['oneTour']),   
    }
</script>