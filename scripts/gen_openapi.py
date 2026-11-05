from pathlib import Path
import mkdocs_gen_files

spec = Path("contracts/openapi.yaml").read_text(encoding="utf-8")
with mkdocs_gen_files.open("api/openapi.yaml", "w") as f:
    f.write(spec)