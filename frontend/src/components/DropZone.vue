<template>
  <div
    @dragenter.prevent="toggleActive"
    @dragleave.prevent="toggleActive"
    @dragover.prevent
    @drop.prevent="toggleActive"
    :class="{ 'bg-gray-200': active }"
    class="dropzone bg-gray-50 border-4 border-gray-300 border-dashed rounded-2xl py-4"
  >
    <div class="py-2">
      <img class="mx-auto w-14 h-14" src="../assets/csv.png" />
    </div>
    <div class="text-gray-500 font-semibold mb-4 text-lg">
      Drag your .csv file here to start uploading.
    </div>
    <div class="mx-auto border-b-2 border-gray-300 leading-[0.1em] w-52">
      <span :class="{ 'bg-gray-200': active }" class="bg-gray-50 px-4 text-gray-400 font-semibold">OR</span>
    </div>
    <div class="mt-10 mb-4">
      <label
        for="dropzoneFile"
        class="bg-primary py-4 px-10 rounded text-white font-semibold cursor-pointer"
        @click="browseFile"
        >Browse File</label
      >
      <input type="file" name="dropzoneFile" class="hidden dropzoneFile" />
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
export default {
  setup() {
    const active = ref(false)
    const toggleActive = () => active.value = !active.value
    const browseFile = () => document.querySelector(".dropzoneFile").click()

    return {
      active,
      toggleActive,
      browseFile
    }
  }
}
</script>
