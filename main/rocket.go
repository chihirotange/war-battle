components {
  id: "rocket"
  component: "/main/rocket.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rocket\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
}
