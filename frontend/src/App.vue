<template>
  <div class="flex flex-col h-screen justify-between">
    <div class="flex-glow">
      <div class="py-32 text-center">
        <div class="max-w-sm mx-auto">
          <h1 class="text-gray-500 font-semibold text-left mb-4 text-lg">Websites Checker</h1>
          <DropZone @drop.prevent="drop" @change="browseFile" />
          <Progress v-if="dropzoneFile" class="mt-8" :file-name="dropzoneFile.name || dropzoneFile" />
          <ResultCard v-if="dropzoneFile" class="mt-8" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DropZone from './components/DropZone.vue'
import Progress from './components/Progress.vue'
import ResultCard from './components/ResultCard.vue'

import { ref } from 'vue';
import axios from 'axios';

export default {
  components: {
    DropZone,
    Progress,
    ResultCard
  },
  setup() {
    let dropzoneFile = ref("");

    const callTest = async() => {
      const test = await axios.get("http://localhost:3000/test");
      console.log("check::", test);
    }

    const drop = (e) => {
      dropzoneFile.value = e.dataTransfer.files[0];
      callTest()
    }

    const browseFile = () => {
      dropzoneFile.value = document.querySelector(".dropzoneFile").files[0];
      callTest()
    }

    return { 
      dropzoneFile,
      drop,
      browseFile
    }
  }
};
</script>
