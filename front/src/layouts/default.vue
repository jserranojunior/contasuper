<template>
  <div >
   <nav >
      <slot name="header"></slot>
    </nav>
    <div class="drawer min-h-screen">
      <input id="my-drawer" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content min-h-screen bg-[#070d18] text-gray-100 font-sans">
      
        <slot name="mainpage"></slot>
      </div>
      <div class="drawer-side">
        <label for="my-drawer" class="drawer-overlay"></label>

        <slot name="sidebar"></slot>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {};
</script>
<style>

</style>
