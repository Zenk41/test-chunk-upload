<template>
  <div>
    <div class="border-solid border-2 border-green-400 m-auto h-fit w-fit p-4">
      <h1>Upload</h1>
      <UForm :validate="validate" :state="state" class="space-y-4" @submit="onSubmit">
        <!-- File input with @change event -->
        <UInput type="file" size="sm" icon="i-heroicons-folder" @change="onSelectedFile" />
        <UButton type="submit">Submit</UButton>
      </UForm>
      <h3>Chunk created:</h3>
      <!-- Display chunks created on the client -->
      <div v-for="(item, index) in chunksClient" :key="index"
        class="border border-green-300 border-solid p-1 rounded-xl my-1">
        Chunk {{ index + 1 }}: {{ item }}
      </div>
      <h3>Chunk needed on the server:</h3>
      <div v-for="(item, index) in chunksServer" :key="index"
        class="border border-blue-300 border-solid p-1 rounded-xl my-1">
        Chunk {{ item }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const chunksClient = ref<number[]>([]);
const chunksServer = ref<number[]>([]);

// File reference
const file = ref<File | null>(null);

// Function to handle file selection
function onSelectedFile(event: FileList) {
  console.log("FileList received:", event);

  if (event.length > 0) {
    file.value = event[0]; // Get the first selected file
    console.log("Selected File:", file.value);

    // Calculate chunk size and create chunks
    calculateChunkSize(file.value.size);
    createChunks(file.value.size);
    
  } else {
    console.warn("No files in FileList");
  }
}

// Function to calculate chunk size
const chunkSize = ref<number>(0);
function calculateChunkSize(fileSize: number) {
  const MAX_CHUNKS = 1000;
  const MIN_CHUNK_SIZE = 1024 * 1024; // 1 MB

  // Calculate optimal chunk size
  const calculatedSize = Math.ceil(fileSize / MAX_CHUNKS);
  chunkSize.value = Math.max(calculatedSize, MIN_CHUNK_SIZE);
  console.log("Calculated chunk size:", chunkSize.value);
}



// Function to create chunks
function createChunks(fileSize: number) {
  chunksClient.value = []; // Reset client chunks

  const totalChunks = Math.ceil(fileSize / chunkSize.value);
  for (let i = 0; i < totalChunks; i++) {
    chunksClient.value.push(i + 1); // Add chunk number
  }

  console.log("Chunks created on the client:", chunksClient.value);

  // Simulate server response (example)
  chunksServer.value = Array.from({ length: totalChunks }, (_, i) => i + 1);
}

// Form submission logic
function onSubmit() {
  console.log("Form submitted with file:", file.value);
  console.log("Chunks to upload:", chunksClient.value);
}

// Dummy form validation and state
const validate = () => true; // Replace with actual validation logic
const state = ref({});

</script>
