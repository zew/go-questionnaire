# pip install pandas pyreadstat
from   pathlib import Path
import pandas as pd
import pyreadstat



def main():
    srcPth   = Path("tmp-bruttostichprobe.dta")
    mixinPth = srcPth.with_name("mixin.csv")
    dstPth   = srcPth.with_name("invitation.csv")

    try:

        df1, meta = pyreadstat.read_dta(str(srcPth), encoding="utf-8")

        # if False:
        if True:
            # print(meta)
            # print(df.head)
            df1 = pd.DataFrame(df1)
            print(df1.describe)
            # 1.443.172 Mio rows

        # renaming
        df1 = df1.rename(columns={"email_firma": "email"})  

        # converting ID to int - avoiding  6434.0
        df1["ID"] = pd.to_numeric(df1["ID"], errors="raise").astype("Int64").astype(str)

        # fixing known email typo without regex interpretation
        df1["email"] = df1["email"].where(df1["email"].notna(), "").astype(str).str.replace("uestarchitekten,de", "uestarchitekten.de", regex=False)
        df1["email"] = df1["email"].where(df1["email"].notna(), "").astype(str).str.replace("info@.hs-umspannwerke.de", "info@hs-umspannwerke.de", regex=False)

        df2 = pd.read_csv(mixinPth, sep="\t", dtype=str, encoding="utf-8", keep_default_na=False)
        if "userid" not in df2.columns:
            raise ValueError("missing userid column in mixin.csv")
        if "link" not in df2.columns:
            raise ValueError("missing link column in mixin.csv")


        # limiting both inputs
        cntLimit = 16166
        cntLimit = 31867
        df1      = df1.head(cntLimit).copy()
        df2      = df2.head(cntLimit).copy()
        ln1      = len(df1)
        ln2      = len(df2)
        if ln1 != ln2:
            raise ValueError(f"row count mismatch: dta={ln1}, mixin={ln2}")


        # inserting mixin values as second and third CSV columns
        df1.insert(1, "userid", df2["userid"].to_list())
        df1.insert(2, "link",   df2["link"  ].to_list())

        df1.to_csv(dstPth, index=False, encoding="utf-8", sep=";")



    except Exception as exc:
        print(exc)
        print("convertStataToCsv")




if __name__ == "__main__":
    main()
