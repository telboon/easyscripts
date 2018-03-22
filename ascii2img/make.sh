cd src
zip -r ascii2png.zip ./*
echo '#!/usr/bin/env python3' | cat - ascii2png.zip > ascii2png

chmod +x ascii2png
mv ascii2png ../
