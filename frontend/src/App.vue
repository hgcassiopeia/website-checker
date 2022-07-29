<template>
  <div class="flex flex-col h-screen justify-between">
    <div class="flex-glow">
      <div class="py-32 text-center">
        <div class="max-w-sm mx-auto">
          <div
            v-if="errorMessage"
            class="w-full p-4 mb-4 bg-red-200 rounded-lg"
          >
            {{ errorMessage }}
          </div>
          <h1 class="text-gray-1 font-semibold text-left mb-4 text-lg">
            Websites Checker
          </h1>
          <DropZone @drop.prevent="drop" @change="browseFile" />
          <Progress
            v-if="uploading"
            class="mt-8"
            :file-name="dropzoneFile.name || dropzoneFile"
            :uploading.sync="progressPercent"
          />
          <ResultCard
            v-if="!uploading && progressPercent == '100'"
            class="mt-8"
            :counter-up="counterUp"
            :counter-down="counterDown"
            :total-links="totalLinks"
            :process-time="processTime"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DropZone from "./components/DropZone.vue";
import Progress from "./components/Progress.vue";
import ResultCard from "./components/ResultCard.vue";

import { ref } from "vue";
import axios from "axios";

export default {
  components: {
    DropZone,
    Progress,
    ResultCard,
  },
  setup() {
    let dropzoneFile = ref("");
    let errorMessage = ref("");
    let uploading = ref(false);
    let progressPercent = ref(0);
    let counterUp = ref(0);
    let counterDown = ref(0);
    let totalLinks = ref(0);
    let processTime = ref("");

    const convertMsToTime = (milliseconds) => {
      let seconds = Math.floor(milliseconds / 1000);
      let minutes = Math.floor(seconds / 60);
      let hours = Math.floor(minutes / 60);
      console.log("test: ", seconds, milliseconds)
      seconds = (seconds % 60) == 0 ? milliseconds + " milliseconds" : (seconds % 60).toString() + " seconds";
      minutes = (minutes % 60) == 0 ? "" : (minutes % 60).toString() + " minutes and ";
      hours = (hours % 24) == 0 ? "" : (hours % 24).toString() + " hours and ";

      return `(Used ${hours} ${minutes} ${seconds})`;
    };

    const callUpload = async (file) => {
      const formData = new FormData();
      formData.append("file", file);

      try {
        uploading.value = true;
        const res = await axios.post("http://localhost:3000/upload", formData, {
          onUploadProgress: (e) =>
            (progressPercent.value = Math.round((e.loaded * 100) / e.total)),
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        errorMessage.value = "";

        if (res.data) {
          const { counter_up, counter_down, total_counter, process_time } =
            res.data;
          counterUp.value = counter_up;
          counterDown.value = counter_down;
          totalLinks.value = total_counter;
          processTime.value = convertMsToTime(process_time);
        }
      } catch (error) {
        console.error(error);
        uploading.value = true;
        errorMessage.value = error.response.data.error;
        setTimeout(() => (errorMessage.value = ""), 5000);

        counterUp.value = 0;
        counterDown.value = 0;
        totalLinks.value = 0;
        processTime.value = "";
      } finally {
        uploading.value = false;
      }
    };

    const drop = (e) => {
      dropzoneFile.value = e.dataTransfer.files[0];
      callUpload(dropzoneFile.value);
    };

    const browseFile = () => {
      dropzoneFile.value = document.querySelector(".dropzoneFile").files[0];
      callUpload(dropzoneFile.value);
    };

    return {
      dropzoneFile,
      errorMessage,
      drop,
      browseFile,
      uploading,
      progressPercent,
      counterUp,
      counterDown,
      processTime,
      totalLinks
    };
  },
};
</script>
