<template>
  <div
    class="flex items-center w-full h-auto rounded-xl bg-white drop-shadow-[0_6px_6px_rgba(0,0,0,0.25)] p-4"
  >
    <div>
      <img class="mx-auto w-12 h-10" src="../assets/csv.png" />
    </div>
    <div class="flex flex-col text-left px-2 w-full">
      <div class="flex justify-between text-xs text-gray-400 mb-1">
        <div>{{fileName}}</div>
        <div class="font-medium">45%</div>
      </div>
      <div class="w-full h-4 bg-gray-200 rounded-full dark:bg-gray-700">
        <div
          class="h-4 bg-primary rounded-full dark:bg-primary"
          style="width: 45%"
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    fileName: {
      type: String,
      required: true
    }
  },
}
</script>