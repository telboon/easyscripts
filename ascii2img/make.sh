pex --python=python3 . -c __main__.py -o ascii2img
echo '#!/usr/bin/env python3' | cat - ascii2img > ascii2img.new
rm ascii2img
mv ascii2img.new ascii2img
chmod +x ascii2img
