<template>
    <nav class="navbar navbar-default" role="navigation">
        <div class="navbar-header">
            <router-link :to="{ name: 'admin'}"><a class="navbar-brand" href="#">Адмін панель</a></router-link>   
        </div>

        <ul class=" navbar-right hr">
            <li><a href="/">Тур агенство</a></li>
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

    ul.hr {
        margin: 10px 0; /* Обнуляем значение отступов */
        padding: 4px; /* Значение полей */
    }
    ul.hr li {
        display: inline; /* Отображать как строчный элемент */
        margin-right: 10px; /* Отступ слева */
        padding: 3px; /* Поля вокруг текста */
    }
</style>