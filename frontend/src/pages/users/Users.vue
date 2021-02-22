<template>
  <section class="flex flex-col items-center justify-center"> 
    <div class="bg-white  w-8/12 m-4 rounded-md shadow flex flex-col p-2">
      <span class="text-center p-4 text-xl uppercase"> Gerenciamento de Usuarios - BB</span>
      <div class="w-full mb-4"> 
        <button class="bg-green-500 text-white py-1 px-2 rounded-md hover:bg-green-700 focus:outline-none" @click="openCreateModal"> + NOVO USUÁRIO </button>
      </div>
      <div class="w-full"> 
        <table v-if="data.length" class="w-full text-md bg-gray-200 shadow-md rounded mb-4">
            <thead>
              <th class="text-left p-3 px-5">Nome</th>
              <th class="text-left p-3 px-5">Sobrenome</th>
              <th class="text-left p-3 px-5">Username</th>
              <th class="text-left w-20 p-3 px-5"></th>
            </thead>
            <tbody>
              <tr v-for="(d, i) in data" :key="d.id" class="border-b hover:bg-orange-100 bg-gray-100">
                <td  class="p-3 px-5">{{ d.name }}</td>
                <td class="p-3 px-5">{{ d.lastname }}</td>
                <td class="p-3 px-5">{{ d.username }}</td>
                <td class="p-3 px-5 flex justify-end w-20">
                  <button 
                    type="button" 
                    class="mr-3 text-sm bg-blue-500 hover:bg-blue-700 text-white py-1 px-2 rounded focus:outline-none focus:shadow-outline" @click="openEditModal(d)">Editar</button>
                  <button 
                    type="button"
                    class="text-sm bg-red-500 hover:bg-red-700 text-white py-1 px-2 rounded focus:outline-none focus:shadow-outline" @click="removeUser(d, i)">Excluir</button>
                </td>
              </tr>
            </tbody>
        </table>
        <div v-else class="flex justify-items-center  bg-gray-100"> 
          <span class="text-center  w-full p-4">Nenhum usuario cadastrado</span>
        </div>
      </div>
    </div>

    <modal name="userModal"  width="300px" height="auto">
      <section class="flex p-4 flex-col justify-center items-center mb-4">
        <UsersForm  :edit="edit" :data="dataSelectedToEdit" @onAdd="addUser"  @onUpdate="updateUser"   @closeForm="closeModal"/>
      </section>
    
    </modal>

  </section>
</template>

<script>
import UsersForm from './UsersForm'
import axios from 'axios'
import { Token } from '../../utils/token'

export default {
  name: "Users",
  components: {
    UsersForm
  },
  data() {
    return {
      edit: false,
      data: [],
      dataSelectedToEdit: null
    }
  },
  async created() {
    try {
       const request = await axios.request({
      url: "http://localhost:5001/api/v1/users",
      method: "GET",
       headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${Token.getToken()}`
          }
    })

    this.data = request.data
    console.log("REsult!!", request.data)
    } catch (error) {
      console.log("errror", error)
    }
  },
  methods: {
    openEditModal(user) {
      this.dataSelectedToEdit = user
      this.edit = true
      this.$modal.show("userModal")
    },
    openCreateModal() {
      this.dataSelectedToEdit =  { name: '', lastname: '', username: ''}
      this.edit = false
      this.$modal.show("userModal")
    },
    addUser(user){
      this.data.push(user)
      this.data = this.data.sort((x, y) => x.name > y.name ? 1 : x.name < y.name ? -1 : 0)
      this.$toastr.s("Usuario adicionado com sucesso");
    },
    updateUser(user){
      let findUser = this.data.findIndex(x => x.id === user.id)
      if (findUser === -1) {
        return 
      }
      this.$set(this.data, findUser, user)
      this.$toastr.s("Usuario atualizado com sucesso");

    },
    async removeUser(user, index){
      try {
              await axios.request({
        method: "DELETE",
        url: `http://localhost:5001/api/v1/users/${user.id}`,
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${Token.getToken()}`
          }
      })

      this.data.splice(index, 1)
      this.$toastr.s("Usuario removido com sucesso");

      }
      catch (err) {
              this.$toastr.s("Erro ao remover usuario");

        console.log(err)
      }

    },
     closeModal(){
      this.$modal.hide("userModal")
    }
   
  }
}
</script>