import pysolr

def get_patchgroup(url='url', searchfor='hostname'):
    solr = pysolr.Solr(url)
    results = solr.search('minion_id:{0}'.format(searchfor))
    if len(results) > 1 :
        return "Narrow the search. More than 1 minion_id found."

    for result in results:
        return result['patchgroup']

if __name__ == '__main__':
    print(get_patchgroup("http://solr/index", "search_term*"))
