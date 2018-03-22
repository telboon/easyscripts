from distutils.core import setup
setup(name='ascii2img',
    version='1.0',
    scripts=['./src/__main__.py'],
    packages=['PIL'],
    package_dir={'': 'src'},
    package_data={'PIL': ['*']},
    )

