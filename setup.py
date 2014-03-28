from setuptools import setup

description = 'PySelecta'
long_desc = open('README.rst').read()

setup(
    name='pyselecta',
    version='0.1.0',
    url='https://github.com/honza/pyselecta',
    install_requires=[],
    description=description,
    long_description=long_desc,
    author='Honza Pokorny',
    author_email='me@honza.ca',
    maintainer='Honza Pokorny',
    maintainer_email='me@honza.ca',
    packages=[],
    include_package_data=True,
    scripts=['pyselecta']
)
