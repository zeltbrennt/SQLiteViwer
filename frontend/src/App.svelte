<script lang="ts">
  import { Events } from "@wailsio/runtime";

  let tableNames: string[] = ["tab 1", "tab 2"];

  Events.On("dbLoaded", (tn: any) => {
    tableNames = tn.data[0];
    console.log(tableNames);
    console.log(tn.data);
  });
</script>

<h1 class="text-2xl">SQLite Viewer</h1>
<div class="tabs tabs-lift tabs-xl">
  {#if tableNames}
    {#each tableNames as db}
      <input type="radio" class="tab" name="db_selected" aria-label={db} />
      <div class="tab-content bg-base-100 border-base-300 p-6">{db}</div>
    {/each}
  {/if}
</div>
