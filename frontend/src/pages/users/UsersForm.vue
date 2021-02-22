<template>
  <section> 
    <p class="text-center uppercase text-lg  border-b-2"> {{ edit ? "Editar" : "Criar"}} usuário</p>
    <FormulateForm name="userForm" class="flex flex-col" values="values" @submit="onSubmit">
      <FormInput label="nome" v-model="data.name" name="name" validation="^required" placeholder="insira um nome" />
      <FormInput label="sobrenome"  v-model="data.lastname" name="lastname" validation="^required" placeholder="insira um sobrenome" />
      <FormInput label="username"  v-model="data.username" name="username" validation="^required" placeholder="insira um username" />
      <FormInput type="password" v-if="!edit"  label="senha" name="password" validation="^required" placeholder="insira uma senha" />
        <FormulateInput
              type="submit"
              :label="edit? 'SALVAR': 'ENVIAR'"
              input-class="bg-blue-500 text-center mt-2 uppercase w-full rounded text-white py-2 px-2 focus:outline-none focus:shadow-outline hover:bg-blue-700 "
            />
        <button type="button" class="bg-red-500 text-center mt-2 uppercase w-full rounded text-white py-2 px-2 focus:outline-none focus:shadow-outline hover:bg-red-700 " @click="closeForm"> Fechar </button>
    </FormulateForm>
  </section>
</template>

<script>
import FormInput from '../../components/forms/Form-input'
import { Token } from '../../utils/token'
import axios from "axios"

export default {
  name: "UserForm",
  components: {
    FormInput
  },
  props: {
    edit: {
      default: false,
    },
    data: {
      default: () =>  ({ name: '', lastname: '', username: ''})
    }
  },
  methods: {
   async onSubmit(data) {
     try {
        const result = await axios.request({
          url: `http://localhost:5001/api/v1/users${this.edit?`/${this.data.id}`:''}`,
          method: this.edit ? "PUT" : "POST",
          data,
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${Token.getToken()}`}
            
          })
        this.edit ? this.$emit("onUpdate", {...data, id: this.data.id }) : this.$emit("onAdd",  result.data)
        !this.edit && this.$formulate.reset("userForm")
      } catch (error) {
       this.$toastr.e(`Ocorreu um erro ao  ${this.edit? 'editar': 'criar'} o usuario`);        
      }
    },
    closeForm() {
      this.$emit("closeForm")
    }
  }
}
</script>