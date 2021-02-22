<template>
    <div class="flex justify-center items-center w-screen h-screen bg-gray-300">
      <div class="flex flex-col bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 mx-2">
        <div class="flex items-center justify-center border-b-2  mb-2">
              <span class="text-center mb-3 text-2xl uppercase tracking-wide  "> LOGIN </span>
          </div>
          <FormulateForm values="values" @submit="onSubmit">
                <FormInput  label="Username" name="username" validation="^required" placeholder="username" />
                <FormInput type="password" label="Senha" name="password" validation="^required" placeholder="password" />
              <div class="w-full mb-1">
                  <FormulateInput
                  type="submit"
                  label="LOGIN"
                  input-class="bg-blue-500 text-center mt-2 uppercase w-full rounded text-white py-2 px-2 focus:outline-none focus:shadow-outline hover:bg-blue-700 "
                />
                <button type="button" @click="openUserForm" class="text-center mt-2 uppercase w-full rounde py-2 px-2  focus:outline-none focus:shadow-outline  hover:underline">
                  Criar conta
                </button>
              </div>
            </FormulateForm>
      </div>


    <modal name="userModal"  width="300px" height="auto">
      <section class="flex p-4 flex-col justify-center items-center mb-4">
        <UsersForm @closeForm="closeModal"/>
      </section>
    
    </modal>
    </div>

</template>
<script>
import FormInput from '../../components/forms/Form-input'
import axios from "axios"
import { Token } from '../../utils/token'
import UsersForm  from '../users/UsersForm'
export default {
  name: "LoginVue",
  components: {
    FormInput,
    UsersForm
  },
  methods:{
   async onSubmit(data){
     try {
      const response = await axios.request({
        url: "http://localhost:5001/api/v1/login",
        method: "POST",
        data: {
          username: data.username,
          password: data.password
        },
        headers: {
          "Content-type": "application/json"
        }
      })
       Token.setToken(response.data)
       this.$router.push("/users")
     } catch(error) {
        this.$toastr.e("Usuario ou senha invalidos");
     }

    },
    openUserForm() {
      this.$modal.show("userModal")
    },
     closeModal(){
      this.$modal.hide("userModal")
    }
    
  }

}
</script>