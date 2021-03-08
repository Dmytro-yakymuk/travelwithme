<template>
    <main id="main">
        <!-- top information area -->
        <div class="inner-top">
            <div class="container">
                <h1 class="inner-main-heading">Оформлення замовлення</h1>
                <!-- breadcrumb -->
                <nav class="breadcrumbs">
                    <ul>
                        <li><a href="#">Home</a></li>
                        <li><a href="#">pages</a></li>
                        <li><span>Checkout</span></li>
                    </ul>
                </nav>
            </div>
        </div>
        <div class="inner-main common-spacing container">
            
            <form class="booking-form" action="#">
                <div class="row">
                    <div class="col-md-10">
                        <div class="form-holder">
                            <h2 class="small-size">Додаткові нотатки</h2>
                             <div class="row">
                                <div class="col-md-4">
                                    <div class="hold">
                                        <label for="phn">Номер телефону*</label>
                                        <input type="text" id="phn" class="form-control">
                                    </div>
                                </div>
                            </div>
                            <div class="wrap">
                                <div class="hold">
                                    <label for="txt">Ваш коментарій</label>
                                    <textarea id="txt" class="form-control" placeholder="Будь ласка, введіть сюди будь-яку додаткову інформацію. "></textarea>
                                </div>
                            </div>
                            <div class="order-block">
                                <h2 class="small-size">Попередній перегляд замовлення</h2>
                                <div class="wrap">
                                    <table class="product-table">
                                        <thead>
                                            <tr>
                                                <th>Вибрані тури</th>
                                                <th>Загальна сума</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td>
                                                    <span class="title">Everest Trekking &emsp;<span class="amount">x&emsp; 3</span></span>
                                                    <time datetime="2016-09-29">Booking Date:&emsp; 14th Jan 2016</time>
                                                </td>
                                                <td>
                                                    <span class="amount">2,999 грн</span>
                                                </td>
                                            </tr>

                                        </tbody>
                                        <tfoot>
                                            <tr>
                                                <td>Загальна сума</td>
                                                <td>4,988 грн</td>
                                            </tr>
                                        </tfoot>
                                    </table>
                                    <button type="submit" class="btn btn-default">Підтвердьте бронювання</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </main>
</template>