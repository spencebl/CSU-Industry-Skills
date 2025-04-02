# CSU-Industry-Skills

Research Project for Comparing CSU and Industry Skills

## Running DPR notebook
DPR_for_gaps.ipynb can be opened using VS Code or Jupyter Notebook software.

NOTE: ONLY DATA PREPROCESSING IS FUNCTIONAL. DPR TRAINING IS STILL BEING CODED.

It is recommended to create a conda environment for the notebook, using Python 3.11.

Be sure to have `descriptions.txt` and `all-csu-codes.csv` within the folder `dpr-model/data/` when executing a fresh training session

### Package Installationspip install pandas

Install necessary packages: 

- `pip`: 
    - `transformers`
    - `wandb`
    - `pandas`
    - `datasets`

- `conda`
    - `scikit-learn`
    - Install Pytorch versioning that matches your device

### Weights and Biases Set-up

Create or log into an account with [Weights & Biases](wandb.ai), to retrieve an API key.
 - Place API key in a txt file named `WANDB_API_KEY.txt`, located at `./dpr-model/WANDB_API_KEY.txt`.
 - Change the `entity` parameter of wandb.init(), within the training function to match your wandb account entity name.

