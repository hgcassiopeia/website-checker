<template>
  <div
    class="flex flex-col w-full h-auto rounded-xl bg-white drop-shadow-[0_6px_6px_rgba(0,0,0,0.25)] p-6"
  >
    <div id="result-header" class="text-left">
        <h1 class="font-medium text-gray-1 text-xl">Total {{totalLinks}} Websites</h1>
        <h3 class="font-medium text-gray-1 text-sm">{{processTime}}</h3>
    </div>
    <div id="result-detail" class="flex justify-around mt-6">
        <div class="bg-green-1 rounded-lg w-20 h-20 p-2">
            <div class="text-white font-medium text-left text-xs">UP</div>
            <div class="text-white font-semibold text-4xl">{{counterUp}}</div>
        </div>
        <div class="bg-gray-4 rounded-lg w-20 h-20 p-2">
            <div class="text-gray-1 font-medium text-left text-xs">Down</div>
            <div class="text-gray-1 font-semibold text-4xl">{{counterDown}}</div>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    counterUp: {
      type: Number,
      required: true,
      default: 0
    },
    counterDown: {
      type: Number,
      required: true,
      default: 0
    },
    totalLinks: {
      type: Number,
      required: true,
      default: 0
    },
    processTime: {
      type: String,
      required: true,
      default: "Not processed"
    }
  },
}
</script>
