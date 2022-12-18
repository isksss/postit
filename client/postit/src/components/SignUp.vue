<template>
    <h2>SignUp</h2>
    <div>
        <input type="email" name="email" id="user-email"><br>
        <input type="password" name="password" id="user-password"><br>
        <!-- <input type="password" name="password" id="user-password2"><br> -->
        <button type="submit" v-on:click="onclick">SignUp</button>
    </div>
</template>

<script>
// import { STRING_UNARY_OPERATORS } from '@babel/types';
    const signupurl = "http://localhost:8000/v1/register"
    export default {
        name: 'SignUp',

        methods:{
            onclick(){
                // console.log($event.timestamp)
                var email = document.getElementById("user-email").value
                console.log(email)

                var password1 = document.getElementById("user-password").value
                console.log(password1)
                // var password2 = document.getElementById("user-password2").value
                // if(password1 != password2){
                //     console.error("パスワードが異なるゾ！")
                // }

                var info = {
                    "email": email,
                    "password": password1 
                }

                var data = {
                    methods: "POST",
                    Headers :{
                        "Content-Type": "application/json"
                    },

                    body: JSON.stringify(info),
                }

                fetch(signupurl, data)
                .then(
                    (res) => {
                        if (!res.ok){
                            console.error("error: not ok...")
                        }
                        console.log("ok!")
                        return res.json()
                    }
                )
                .then(
                    data => {
                        console.log(data)
                    }
                )
                .catch(
                    error =>{
                        console.error(error)
                    }
                )
            }
        }
    }
</script>

