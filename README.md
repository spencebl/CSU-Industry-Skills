# CSU-Industry-Skills

Research Project for Comparing CSU and Industry Skills

## Running BERT Model

gaps_bert.ipynb can be opened using VS Code or Jupyter Notebook software.

It is recommended to create a conda environment for the notebook, using Python 3.11

### Data Processing

If you have the file `.\bert-model\data\bert_training_data.csv`, skip this section.

If you don't have the file `.\bert-model\data\bert_training_data.csv`, uncomment the code in cells 2, 3, 8, and 9 to create the file.

- Be sure to have `descriptions.txt` and `all-csu-codes.csv` within the folder `bert-model/data/` when executing a fresh training session.
- Once the training data is created, comment out those cells again.

### Package Installationspip install pandas

Install necessary packages:

`pip`:

- `wandb`
- `pandas`
- `sentence-transformers`

`conda`:

- `scikit-learn`
- Install Pytorch versioning that matches your device

### Weights and Biases Set-up

Create or log into an account with [Weights & Biases](wandb.ai), to retrieve an API key.

- Place API key in a txt file named `WANDB_API_KEY.txt`, located at `./dpr-model/WANDB_API_KEY.txt`.
- Change the `entity` parameter of wandb.init(), within the training function to match your wandb account entity name.
