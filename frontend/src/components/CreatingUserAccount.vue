<template>
  <div>
    <label for="name">Name<input></label>
    <label for="birthday">Birthday<input></label>
  </div>
  <button type="submit" @onclick=""></button>
</template>

<script setup lang="ts">
const response = await fetch('http://localhost:8080/api/v1/userdata', {
  method: 'POST',
  headers: {
    'Content-Type': 'multipart/form-data'
  }
});
if (!response.ok) {
  throw new Error(response.statusText);
}
const data = await response.json();

</script>