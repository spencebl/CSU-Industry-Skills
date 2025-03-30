# CSU-Industry-Skills

Research Project for Comparing CSU and Industry Skills

## Running DPR notebook

DPR_for_gaps.ipynb can be opened using VS Code or Jupyter Notebook software.

NOTE: ONLY DATA PREPROCESSING IS FUNCTIONAL. DPR TRAINING IS STILL BEING CODED.

- It is recommended to create a conda environment for the notebook.

### Package Installations

- Install necessary packages: `pip install transformers, wandb, pandas, datasets`

- install Pytorch versioning that matches your device via conda 

- The notebook was executed using Python v3.13.2 within the conda environemnt.

- Be sure to have `descriptions.txt` and `all-csu-codes.csv` within the folder `dpr-model/data/` when executing a fresh training session
