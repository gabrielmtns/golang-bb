<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  inheritAttrs: true,
  name: 'FormInput',
  props: {
    label: String,
    name: String,
    placeholder: String,
    validation: String,
    value: {
      type: [String, Number],
      default: null
    },
    id: String,
    type: {
      type: String,
      default: 'text'
    },
    blockPaste: {
      default: false,
      type: Boolean
    }
  }
})
</script>

<template>

  <div>
    <FormulateInput
      :value="value"
      :type="type"
      :name="name"
      :label="label"
      v-bind="$attrs" 
      label-class="mt-4 block uppercase text-gray-700 text-xs font-bold "
      :validation="validation"
      :validation-messages="{
        required: 'Campo obrigatório',
        email: 'Formato de e-mail invalido.',
        confirm: 'Confirmação de campo invalida. Por favor verifique se digitou corretamente.'
      }"
      :placeholder="placeholder"
      input-class="w-full  px-3 py-2 border border-gray-400 border-box rounded leading-none focus:border-blue-500 outline-none"
      error-class="float-left text-red-500 text-xs italic border-red-500"
      input-hasErrors-class="appearance-none border border-red-500"
    />
  </div>
</template>