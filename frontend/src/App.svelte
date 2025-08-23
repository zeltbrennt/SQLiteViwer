<script lang="ts">
  import { Events } from "@wailsio/runtime";

  let tableNames: string[] = [];

  Events.On("dbLoaded", (tn: any) => {
    tableNames = tn.data[0];
    console.log(tableNames);
    console.log(tn.data[0]);
  });
</script>

<h1>SQLite Viewer</h1>
{#if tableNames}
  {#each tableNames as db}
    <button class="btn">{db}</button>
  {/each}
{/if}
